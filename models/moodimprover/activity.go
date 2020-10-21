package moodimprover

import "gorm.io/gorm"

// Activity is a thing todo which will improve your mood!
type Activity struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	EntryID     uint
}

// New returns a new activity.
func New(title string, description string) *Activity {
	return &Activity{
		Title:       title,
		Description: description,
	}
}
