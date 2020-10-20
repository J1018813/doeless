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
	ID                    int
	DateTime              time.Time                   `json:"dateTime"`
	GratitudeDescriptions []mood.GratitudeDescription `json:"gratitudeDescriptions"`
	Mood                  mood.Mood                   `json:"mood"`
	Activity              moodimprover.Activity       `json:"activity"`
}

// New gives you a new Entry.
func New(id int) *Entry {
	return &Entry{
		ID:       id,
		DateTime: time.Now(),
	}
}
