package store

import (
	"context"
	"sort"

	"github.com/DF-wu/HideReplier/internal/config"
	"github.com/DF-wu/HideReplier/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	client     *mongo.Client
	database   *mongo.Database
	historyCol *mongo.Collection
	counterCol *mongo.Collection
}

// NewMongoStore opens the MongoDB collections used by the Go port.
func NewMongoStore(ctx context.Context, cfg config.Config) (*MongoStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return nil, err
	}

	database := client.Database(cfg.MongoDatabase)

	return &MongoStore{
		client:     client,
		database:   database,
		historyCol: database.Collection("DiscordPostCollection"),
		counterCol: database.Collection("Counter"),
	}, nil
}

// Close disconnects the shared MongoDB client.
func (m *MongoStore) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

// LoadOrCreateCounter loads the serial counter document or creates it on first boot.
func (m *MongoStore) LoadOrCreateCounter(ctx context.Context) (*model.SerialCounter, error) {
	var counters []model.SerialCounter
	cursor, err := m.counterCol.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &counters); err != nil {
		return nil, err
	}

	if len(counters) > 0 {
		return &counters[0], nil
	}

	initial := model.SerialCounter{Counter: 0}
	result, err := m.counterCol.InsertOne(ctx, initial)
	if err != nil {
		return nil, err
	}

	if id, ok := result.InsertedID.(interface{ Hex() string }); ok {
		initial.ID = id.Hex()
	}

	return &initial, nil
}

// SaveCounter replaces the persisted counter document and fails if it matched nothing.
func (m *MongoStore) SaveCounter(ctx context.Context, counter *model.SerialCounter) error {
	result, err := m.counterCol.ReplaceOne(ctx, bson.M{"_id": counter.ID}, counter)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

// InsertHistory stores one successful anonymous message submission.
func (m *MongoStore) InsertHistory(ctx context.Context, data model.StoreData) error {
	_, err := m.historyCol.InsertOne(ctx, data)
	return err
}

// ListHistory returns all stored history rows sorted by serial number.
func (m *MongoStore) ListHistory(ctx context.Context) ([]model.StoreData, error) {
	var history []model.StoreData
	cursor, err := m.historyCol.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &history); err != nil {
		return nil, err
	}

	sort.Slice(history, func(i, j int) bool {
		return history[i].SerialNumber < history[j].SerialNumber
	})

	return history, nil
}
