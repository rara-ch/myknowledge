package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rara-ch/myknowledge.git/internal/database"
)

type state struct {
	db *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	connString := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer db.Close()

	s := &state{
		db: database.New(db),
	}

	cmds := newCLICommands()
	cmds.register("add", addTopicHandler)

	input := os.Args

	if len(input) < 2 {
		fmt.Println("no commands entered")
		os.Exit(1)
	}

	cmd := inputCommand{
		name: input[1],
		args: input[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
