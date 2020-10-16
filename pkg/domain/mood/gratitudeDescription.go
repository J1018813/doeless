package mood

import "time"

// GratitudeDescription describes what a user is gratefull for today.
type GratitudeDescription struct {
	DateTime    time.Time
	Title       string
	Description string
}

// New returns a new Gratitude description.
func New(title, string, description string) *GratitudeDescription {
	return &GratitudeDescription{
		DateTime:    time.Now(),
		Title:       title,
		Description: description,
	}
}
