package container

import (
	"cleverHouse-go-back/config"
	"cleverHouse-go-back/database"
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

type Container struct {
	UserRepo database.UserRepository
}

func New() Container {
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

	fmt.Println(sess)

	var ur = database.NewUserRepository(sess)

	return Container{UserRepo: ur}
}
