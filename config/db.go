package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
func CoonectDb() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")

	client,err:= mongo.Connect(ctx,clientOption)

	if err!=nil{
		log.Fatal("Database connection failed")
	}
	DB = client

}