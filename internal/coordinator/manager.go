package coordinator

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/khosbilegt/llama-drover/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var db *mongo.Database

func Init(database *mongo.Database) {
	db = database
	log.Println("Coordinator initialized with database:", db.Name())
}

func CreateHerd(name string, modelName string) (model.Herd, error) {
	id := uuid.New()

	var herd = model.Herd{
		Name:  name,
		ID:    id.String(),
		Model: modelName,
		Nodes: []model.Node{},
	}
	_, err := db.Collection("herds").InsertOne(context.Background(), herd)
	if err != nil {
		log.Println("Error creating herd:", err)
		return model.Herd{}, err
	}
	return herd, nil
}

func DeleteHerd(herdID string) error {
	log.Println("Deleting herd with ID:", herdID)
	_, err := db.Collection("herds").DeleteOne(context.Background(), bson.M{"id": herdID})
	if err != nil {
		log.Println("Error deleting herd:", err)
		return err
	}
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

func CreateNode(node model.Node) error {
	log.Println("Creating node:", node.Name)
	// Here you would typically insert the node into the database
	return nil
}

func DeleteNode(nodeID string) error {
	log.Println("Deleting node with ID:", nodeID)
	// Here you would typically remove the node from the database
	return nil
}

func GetNode(nodeID string) (model.Node, error) {
	var node model.Node
	err := db.Collection("nodes").FindOne(context.Background(), bson.M{"id": nodeID}).Decode(&node)
	if err != nil {
		log.Println("Error getting node:", err)
		return model.Node{}, err
	}
	return node, nil
}

func ListNodes() ([]model.Node, error) {
	var nodes []model.Node
	cursor, err := db.Collection("nodes").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error listing nodes:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var node model.Node
		if err := cursor.Decode(&node); err != nil {
			log.Println("Error decoding node:", err)
			continue
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
