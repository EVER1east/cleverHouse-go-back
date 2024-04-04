package database

import (
	"cleverHouse-go-back/domain"

	"github.com/upper/db/v4"
)

const UsersTableName = "users"

type user struct {
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Phone      string `db:"phone"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Id         uint64 `db:"id,omitempty"`
}
type UserRepository struct {
	coll db.Collection
	sess db.Session
}

func NewUserRepository(session db.Session) UserRepository {
	return UserRepository{
		coll: session.Collection(UsersTableName),
		sess: session,
	}
}
func (r UserRepository) Save(u domain.User) (domain.User, error) {
	var user = r.mapDomainToModel(u)
	var err = r.coll.InsertReturning(&user)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(user), nil
}

func (r UserRepository) mapDomainToModel(u domain.User) user {
	return user{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Phone:      u.Phone,
		Email:      u.Email,
		Password:   u.Password,
	}
}

func (r UserRepository) mapModelToDomain(u user) domain.User {

	return domain.User{
		Id:         u.Id,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Phone:      u.Phone,
		Email:      u.Email,
		Password:   u.Password,
	}
}
