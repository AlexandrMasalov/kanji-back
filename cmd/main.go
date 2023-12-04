package main

import (
	"github.com/AlexandrMasalov/kanji/internal/connectdb"
	"github.com/AlexandrMasalov/kanji/internal/server"
)

const connectionString = "host=127.0.0.1 port=5432 user=admin dbname=kanji sslmode=disable"


func main() {
	connectdb.ConnectDB(connectionString)
	server.Run()
}
