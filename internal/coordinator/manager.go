package coordinator

import (
	"context"
	"log"

	"github.com/khosbilegt/llama-drover/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var db *mongo.Database

func Init(database *mongo.Database) {
	db = database
	log.Println("Coordinator initialized with database:", db.Name())
}

func CreateHerd(name string) (model.Herd, error) {
	return model.Herd{Name: name}, nil
}

func DeleteHerd(herdID string) error {
	log.Println("Deleting herd with ID:", herdID)
	return nil
}

func GetHerd(herdID string) (model.Herd, error) {
	var herd model.Herd
	err := db.Collection("herds").FindOne(context.Background(), bson.M{"id": herdID}).Decode(&herd)
	if err != nil {
		log.Println("Error getting herd:", err)
		return model.Herd{}, err
	}
	return herd, nil
}

func ListHerds() ([]model.Herd, error) {
	var herds []model.Herd
	cursor, err := db.Collection("herds").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error listing herds:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var herd model.Herd
		if err := cursor.Decode(&herd); err != nil {
			log.Println("Error decoding herd:", err)
			continue
		}
		herds = append(herds, herd)
	}

	return herds, nil
}
