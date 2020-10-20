package user

import (
	"github.com/j1018813/doeless/models/entry"
	"gorm.io/gorm"
)

// User is the kind person using this application.
type User struct {
	gorm.Model
	ID      int
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Entries []*entry.Entry `json:"entries"`
}

// New creates a new user.
func New(id int, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// AddEntry assignes the given entry to the user.
func (u *User) AddEntry(entry *entry.Entry) {
	u.Entries = append(u.Entries, entry)
}
