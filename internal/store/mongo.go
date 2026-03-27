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

func (m *MongoStore) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

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

func (m *MongoStore) InsertHistory(ctx context.Context, data model.StoreData) error {
	_, err := m.historyCol.InsertOne(ctx, data)
	return err
}

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
