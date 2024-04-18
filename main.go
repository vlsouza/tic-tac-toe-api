package main

import (
	"fmt"
	"log"
	"main/match"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
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
