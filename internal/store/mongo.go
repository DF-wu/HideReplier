package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	historyCollection = "DiscordPostCollection"
	counterCollection = "Counter"

	// The app runs on a single shared-CPU machine with a hard connection
	// limit of a few dozen HTTP requests, so a small pool is plenty and keeps
	// idle sockets (and their goroutines/buffers) off the 256MB budget.
	maxPoolSize = 8

	// historyBatchSize bounds how many history rows the driver buffers from
	// the cursor at a time while streaming GET /HideBot/discord.
	historyBatchSize = 200
)

type MongoStore struct {
	client     *mongo.Client
	historyCol *mongo.Collection
	counterCol *mongo.Collection
	counterID  any
}

// NewMongoStore opens the MongoDB collections used by the Go port, locates
// (or creates) the serial counter document and makes sure the history
// collection can be sorted by serial number on the server.
func NewMongoStore(ctx context.Context, cfg config.Config) (*MongoStore, error) {
	clientOpts := options.Client().
		ApplyURI(cfg.MongoURI).
		SetMaxPoolSize(maxPoolSize).
		SetMinPoolSize(0).
		SetMaxConnIdleTime(5 * time.Minute).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	database := client.Database(cfg.MongoDatabase)
	store := &MongoStore{
		client:     client,
		historyCol: database.Collection(historyCollection),
		counterCol: database.Collection(counterCollection),
	}

	bootCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := store.initCounter(bootCtx); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("init serial counter: %w", err)
	}

	// Index creation is best effort: a read-only database user must not stop
	// the service from booting. Without the index Mongo sorts in memory.
	if err := store.ensureHistoryIndex(bootCtx); err != nil {
		log.Printf("warning: could not ensure history index on serialNumber: %v", err)
	}

	return store, nil
}

// Close disconnects the shared MongoDB client.
func (m *MongoStore) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoStore) initCounter(ctx context.Context) error {
	var counter model.SerialCounter
	err := m.counterCol.FindOne(ctx, bson.D{}).Decode(&counter)
	switch {
	case err == nil:
		m.counterID = counter.ID
		return nil
	case err == mongo.ErrNoDocuments:
		result, insertErr := m.counterCol.InsertOne(ctx, model.SerialCounter{Counter: 0})
		if insertErr != nil {
			return insertErr
		}
		m.counterID = result.InsertedID
		return nil
	default:
		return err
	}
}

func (m *MongoStore) ensureHistoryIndex(ctx context.Context) error {
	_, err := m.historyCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "serialNumber", Value: 1}},
	})
	return err
}

// NextSerial atomically reserves the next serial number in MongoDB. This is a
// single round trip and is safe across concurrent requests and multiple
// running instances (e.g. during a rolling deploy).
func (m *MongoStore) NextSerial(ctx context.Context) (int, error) {
	var updated model.SerialCounter
	err := m.counterCol.FindOneAndUpdate(
		ctx,
		bson.M{"_id": m.counterID},
		bson.M{"$inc": bson.M{"counter": 1}},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updated)
	if err != nil {
		return 0, err
	}

	return updated.Counter, nil
}

// InsertHistory stores one successful anonymous message submission.
func (m *MongoStore) InsertHistory(ctx context.Context, data model.StoreData) error {
	_, err := m.historyCol.InsertOne(ctx, data)
	return err
}

// StreamHistory hands history rows to fn in ascending serial-number order.
//
// With limit <= 0 the whole collection is streamed straight from the cursor,
// so memory stays bounded by one batch regardless of collection size. With
// limit > 0 only the newest `limit` rows are fetched (sorted descending on
// the server, then reversed here), which is what a "recent posts" caller wants.
func (m *MongoStore) StreamHistory(ctx context.Context, limit int64, fn func(model.StoreData) error) error {
	if limit > 0 {
		return m.recentHistory(ctx, limit, fn)
	}

	findOpts := options.Find().
		SetSort(bson.D{{Key: "serialNumber", Value: 1}}).
		SetBatchSize(historyBatchSize)

	cursor, err := m.historyCol.Find(ctx, bson.D{}, findOpts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var row model.StoreData
		if err := cursor.Decode(&row); err != nil {
			return err
		}
		if err := fn(row); err != nil {
			return err
		}
	}

	return cursor.Err()
}

func (m *MongoStore) recentHistory(ctx context.Context, limit int64, fn func(model.StoreData) error) error {
	findOpts := options.Find().
		SetSort(bson.D{{Key: "serialNumber", Value: -1}}).
		SetLimit(limit)

	cursor, err := m.historyCol.Find(ctx, bson.D{}, findOpts)
	if err != nil {
		return err
	}

	var rows []model.StoreData
	if err := cursor.All(ctx, &rows); err != nil {
		return err
	}

	for i := len(rows) - 1; i >= 0; i-- {
		if err := fn(rows[i]); err != nil {
			return err
		}
	}

	return nil
}
