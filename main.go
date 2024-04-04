package main

import (
	"cleverHouse-go-back/config/container"
	"cleverHouse-go-back/domain"
	"fmt"
	"log"
)

func main() {
	var cont = container.New()

	var user = domain.User{
		FirstName:  "Vitaliy",
		SecondName: "Obrigaliy",
		Email:      "kostastepan7@gmail.com",
		Phone:      "+0958766135",
		Password:   "123",
	}

	var newUser, err = cont.UserRepo.Save(user)
	if err != nil {
		log.Printf("Errrrror: %s", err)
	}

	fmt.Println(newUser)
}
