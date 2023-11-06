package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/MerchLeti/service/internal/db"
	"github.com/MerchLeti/service/internal/env"
	"github.com/MerchLeti/service/internal/server"
)

const defaultPort = "9000"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	database, err := db.New(ctx)
	if err != nil {
		log.Fatalf("couldn't setup database: %v", err)
	}
	port := env.Get("SERVER_PORT", defaultPort)
	log.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		server.New(database),
	); err != nil {
		log.Fatal(err)
	}
}
