package database

import (
	"github.com/egor71717/SquadTechnologiesTestApp/model"
	"github.com/jinzhu/gorm"
)

//PostgreSQLDB is DB interface implementation
type PostgreSQLDB struct {
	Db *gorm.DB
}

//GetUser gets user with id from db
func (psqldb PostgreSQLDB) GetUser(id int) (model.User, error) {
	var user model.User
	psqldb.Db.First(&user, id)
	return user, nil
}

//GetUserByLogin gets user with login from db
func (psqldb PostgreSQLDB) GetUserByLogin(login string) (*model.User, error) {
	var user model.User
	psqldb.Db.First(&user, "login = ?", login)
	if user.Login == "" {
		return nil, ErrorNotFound
	}
	return &user, nil
}

//GetUsers gets all users from db
func (psqldb PostgreSQLDB) GetUsers() ([]model.User, error) {
	var users []model.User
	psqldb.Db.Find(&users)
	return users, nil
}

//CreateUser creates a new user
func (psqldb PostgreSQLDB) CreateUser(user *model.User) {
	psqldb.Db.Create(user)
}
