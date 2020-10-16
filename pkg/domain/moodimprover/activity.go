package moodimprover

// Activity is a thing todo which will improve your mood!
type Activity struct {
	Title       string
	Description string
}

// New returns a new activity.
func New(title string, description string) *Activity {
	return &Activity{
		Title:       title,
		Description: description,
	}
}
