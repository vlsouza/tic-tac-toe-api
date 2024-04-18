package main

import (
	"context"
	"fmt"
	"log"
	"main/match"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/mux"
)

func main() {
	db := initDBClient()
	initAPI(db)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tic-tac-toe-api is running...")
}

func initAPI(db *dynamodb.Client) {
	//check API health
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health)

	config := match.Config{
		DB:     db,
		Router: router,
	}

	//initiate API
	match.NewAPI(config)

	fmt.Println("API is running at port 8080.")

	port := "8080"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func initDBClient() *dynamodb.Client {
	ctx := context.Background()

	fmt.Println("Connecting to DB........")

	// Carregando configurações padrão
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("eu-north-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	fmt.Println("Connected.")

	// Criando o cliente DynamoDB
	return dynamodb.NewFromConfig(cfg)
}
