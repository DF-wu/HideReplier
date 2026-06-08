package model

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSerialCounterKeepsMongoObjectIDType(t *testing.T) {
	id := primitive.NewObjectID()
	counter := SerialCounter{ID: id, Counter: 42}

	encoded, err := bson.Marshal(counter)
	if err != nil {
		t.Fatalf("marshal counter: %v", err)
	}

	var decoded bson.M
	if err := bson.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unmarshal counter: %v", err)
	}

	got, ok := decoded["_id"].(primitive.ObjectID)
	if !ok {
		t.Fatalf("expected ObjectID _id after BSON round trip, got %T", decoded["_id"])
	}

	if got != id {
		t.Fatalf("expected _id %s, got %s", id.Hex(), got.Hex())
	}
}
