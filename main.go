package main

import (
	"cleverHouse-go-back/config"
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

func main() {
	var config = config.GetConfig()

	var settings = postgresql.ConnectionURL{
		Database: config.DbName,
		Host:     config.DbHost,
		User:     config.DbUser,
		Password: config.DbPassword,
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatalf("Ми всрали з'єднання з базою даних: %s", err)
	}

	defer sess.Close()

	fmt.Println(sess)

}
