package database

import (
	"github.com/j1018813/doeless/models/user"

	"gorm.io/gorm"
)

// UserDatabase contains the methods to interact with the users in the database.
type UserDatabase struct {
	db *gorm.DB
}

// NewUserDatabase gives you a new User Database
func NewUserDatabase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{
		db: db,
	}
}

// GetUsers retrieves all users from the database.
func (u *UserDatabase) GetUsers() ([]user.User, error) {
	var users []user.User
	u.db.Find(&users)
	return users, nil
}

// NewUser persists a new user to the database.
func (u *UserDatabase) NewUser(user *user.User) error {
	u.db.Create(user)
	return nil
}
