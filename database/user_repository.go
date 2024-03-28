package database

import "github.com/upper/db/v4"

type user struct {
	Id         uint64 `db: "id,omitempty"`
	Phone      string `db: "phone"`
	Email      string `db: "email"`
	FirstName  string `db: "first_name"`
	SecondName string `db: "second_name"`
	Password   string `db: "password"`
}

func SaveVitaliy(sess db.Session) error {
	var u = user{
		FirstName:  "Vitaliy",
		SecondName: "Obrigaliy",
		Phone:      "+380958766135",
		Email:      "kostastepan7@gmail.com",
		Password:   "87654321",
	}

	var err = sess.
		Collection("users").
		InsertReturning(&u)

	return err
}
