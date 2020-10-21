package entry

import (
	"time"

	"github.com/j1018813/doeless/models/mood"
	"github.com/j1018813/doeless/models/moodimprover"
	"gorm.io/gorm"
)

// Entry tells you how the user is doing that day.
type Entry struct {
	gorm.Model
	DateTime              time.Time                   `json:"dateTime"`
	GratitudeDescriptions []mood.GratitudeDescription `json:"gratitudeDescriptions"`
	Mood                  mood.Mood                   `json:"mood"`
	Activity              moodimprover.Activity       `json:"activity"`
	UserID                uint
}

// New gives you a new Entry.
func New(id int) *Entry {
	return &Entry{
		DateTime: time.Now(),
	}
}
