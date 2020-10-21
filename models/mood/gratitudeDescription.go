package mood

import (
	"time"

	"gorm.io/gorm"
)

// GratitudeDescription describes what a user is gratefull for today.
type GratitudeDescription struct {
	gorm.Model
	DateTime    time.Time `json:"dateTime"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EntryID     uint
}

// New returns a new Gratitude description.
func New(title, string, description string) *GratitudeDescription {
	return &GratitudeDescription{
		DateTime:    time.Now(),
		Title:       title,
		Description: description,
	}
}
