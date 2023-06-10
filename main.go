package main

import (
	"database/sql"
	"log"

	"github.com/kien6034/golang-BE/api"
	db "github.com/kien6034/golang-BE/db/sqlc"
	"github.com/kien6034/golang-BE/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal(("cannot start server:"), err.Error())
	}
}
