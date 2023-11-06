package db

import (
	"fmt"

	"github.com/MerchLeti/service/internal/env"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "shop"
	password = "shop"
	dbname   = "shop"
)

func DSNFromEnv() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Get("POSTGRES_HOST", host),
		env.Get("POSTGRES_PORT", port),
		env.Get("POSTGRES_USER", user),
		env.Get("POSTGRES_PASSWORD", password),
		env.Get("POSTGRES_DB", dbname),
	)
}
