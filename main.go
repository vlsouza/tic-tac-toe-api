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
	initDBClient()
	initAPI()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tic-tac-toe-api is running...")
}

func initAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health)

	config := match.Config{Router: router}

	match.NewAPI(config)

	fmt.Println("API is running at port 8080.")

	port := "8080"
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func initDBClient() {
	ctx := context.Background()

	// Carregando configurações padrão
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("eu-north-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Criando o cliente DynamoDB
	svc := dynamodb.NewFromConfig(cfg)

	// Chamada de exemplo: Listar tabelas
	result, err := svc.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}

	fmt.Println("Tables:")
	for _, tableName := range result.TableNames {
		fmt.Println(tableName)
	}
}
