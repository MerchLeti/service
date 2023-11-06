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

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	database, err := db.New(ctx)
	if err != nil {
		log.Fatalf("couldn't setup database: %v", err)
	}
	if err := http.ListenAndServe(
		fmt.Sprintf(":%v", env.Get("SERVER_PORT", "9000")),
		server.New(database),
	); err != nil {
		log.Fatal(err)
	}
}
