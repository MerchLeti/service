package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/MerchLeti/service/internal/db"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var migrations embed.FS

func main() {

	if len(os.Args) < 2 {
		usage()
		return
	}

	dbString := db.DSNFromEnv()
	command := os.Args[1]

	if _, disabled := disabledCommands[command]; disabled {
		usage()
		return
	}

	database, err := goose.OpenDBWithDriver("pgx", dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	var arguments []string
	if len(os.Args) > 2 {
		arguments = append(arguments, os.Args[2:]...)
	}

	goose.SetBaseFS(migrations)
	if err := goose.Run(command, database, ".", arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}

}

func usage() {
	fmt.Println("Usage: applydb COMMAND")
	fmt.Println(usageCommands)
}

var disabledCommands = map[string]struct{}{
	"create": {},
}

const usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    fix                  Apply sequential ordering to migrations
    validate             Check migration files without running them`
