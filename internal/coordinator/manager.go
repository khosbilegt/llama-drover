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

func CreateCluster(name string, modelName string) (model.Cluster, error) {
	id := uuid.New()

	var cluster = model.Cluster{
		Name:  name,
		ID:    id.String(),
		Model: modelName,
		Nodes: []model.Node{},
	}
	_, err := db.Collection("clusters").InsertOne(context.Background(), cluster)
	if err != nil {
		log.Println("Error creating cluster:", err)
		return model.Cluster{}, err
	}
	return cluster, nil
}

func DeleteCluster(clusterID string) error {
	log.Println("Deleting cluster with ID:", clusterID)
	_, err := db.Collection("clusters").DeleteOne(context.Background(), bson.M{"id": clusterID})
	if err != nil {
		log.Println("Error deleting cluster:", err)
		return err
	}
	return nil
}

func GetCluster(clusterID string) (model.Cluster, error) {
	var cluster model.Cluster
	err := db.Collection("clusters").FindOne(context.Background(), bson.M{"id": clusterID}).Decode(&cluster)
	if err != nil {
		log.Println("Error getting cluster:", err)
		return model.Cluster{}, err
	}
	return cluster, nil
}

func ListClusters() ([]model.Cluster, error) {
	var clusters []model.Cluster
	cursor, err := db.Collection("clusters").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error listing clusters:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var cluster model.Cluster
		if err := cursor.Decode(&cluster); err != nil {
			log.Println("Error decoding cluster:", err)
			continue
		}
		clusters = append(clusters, cluster)
	}

	return clusters, nil
}

func CreateNode(nodeInput model.Node) (model.Node, error) {
	node := model.Node{
		ID:        uuid.New().String(),
		Name:      nodeInput.Name,
		IPAddress: nodeInput.IPAddress,
		Port:      nodeInput.Port,
		Status:    "INITIALIZING",
	}

	_, err := db.Collection("nodes").InsertOne(context.Background(), node)
	if err != nil {
		log.Println("Error creating node:", err)
		return model.Node{}, err
	}
	return node, nil
}

func DeleteNode(nodeID string) error {
	_, err := db.Collection("nodes").DeleteOne(context.Background(), bson.M{"id": nodeID})
	if err != nil {
		log.Println("Error deleting node:", err)
		return err
	}
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
