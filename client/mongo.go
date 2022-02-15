package client

import (
	"log"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/LNMMusic/fce/config"
)

// INSTANCE
type MongoInstance struct {
	Client		*mongo.Client
	Db			*mongo.Database
	Cl			*mongo.Collection
}
var Mongo MongoInstance


// CONNECTION
func (m *MongoInstance) ConnectClient() {
	// URI [Client]
	client, err := mongo.NewClient(options.Client().ApplyURI(config.EnvGet("MONGODB_URI")))
	if err != nil {
		log.Fatalf("Failed to Create New Client Mongo-DB -> %v", err)
	}
	m.Client = client

	// URI [Connect]
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx); if err != nil {
		log.Fatalf("Failed to Connect to Mongo-DB -> %v", err)
	}

	log.Println("Succeed to Connect to Mongo-DB")
}
func (m *MongoInstance) DisconnectClient() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := m.Client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to Disconnect from MONGODB -> %v", err)
	}
}

// DB + CL
func (m *MongoInstance) ConnectDB(name string) {
	m.Db = m.Client.Database(name)
}
func (m *MongoInstance) ConnectCL(name string) {
	m.Cl = m.Db.Collection(name)
}