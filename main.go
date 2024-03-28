package main

import (
	"cleverHouse-go-back/config"

	"github.com/upper/db/v4/adapter/postgresql"
)

func main() {
	var config = config.GetConfig()

	var settings = postgresql.ConnectionURL{
		Database: config.DbName,
		Host: config.DbHost,
		User: config.DbUser,
		Password: config.DbPassword,
	}
}
