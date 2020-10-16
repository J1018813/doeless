package entry

import (
	"time"

	"github.com/j1018813/doeless/pkg/domain/mood"
	"github.com/j1018813/doeless/pkg/domain/moodimprover"
)

// Entry tells you how the user is doing that day.
type Entry struct {
	ID                    int
	DateTime              time.Time
	GratitudeDescriptions []mood.GratitudeDescription
	Mood                  mood.Mood
	Activity              moodimprover.Activity
}

// New gives you a new Entry.
func New(id int) *Entry {
	return &Entry{
		ID:       id,
		DateTime: time.Now(),
	}
}
