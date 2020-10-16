package main

import (
	"github.com/j1018813/doeless/pkg/domain/entry"
	"github.com/j1018813/doeless/pkg/domain/user"
)

func main() {
	me := user.New(1, "Jorn", "test@test.com")

	e := entry.New(1)

	me.AddEntry(e)

	e.Mood = 4

	println("mood: ", me.Entries[0].Mood)

}
