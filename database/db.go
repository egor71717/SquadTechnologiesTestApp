package database

import (
	"errors"

	"github.com/egor71717/SquadTechnologiesTestApp/model"
)

//ErrorNotFound is a not found error
var ErrorNotFound = errors.New("not found")

// DB is the interface for application storage
type DB interface {

	// returns the User for the given id, ErrNotFound if the key doesn't exist,
	// or another error if the get failed
	GetUser(id int) (model.User, error)
	// returns the User for the given login, ErrNotFound if the key doesn't exist,
	// or another error if the get failed
	GetUserByLogin(login string) (*model.User, error)
	// returns all Users, or an error if the get failed
	GetUsers() ([]model.User, error)
	// Create creates a new user
	CreateUser(*model.User)
}
