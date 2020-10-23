package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/j1018813/doeless/database"
	"github.com/j1018813/doeless/handlers"
	"github.com/j1018813/doeless/models/entry"
	"github.com/j1018813/doeless/models/mood"
	"github.com/j1018813/doeless/models/moodimprover"
	"github.com/j1018813/doeless/models/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("database connection successfuly opened")

	database.Db.AutoMigrate(
		&user.User{},
		&entry.Entry{},
		&moodimprover.Activity{},
		&mood.GratitudeDescription{},
	)
}

func setupRoutes(app *fiber.App) {

	userDatabase := database.NewUserDatabase(database.Db)
	userHandlers := handlers.NewUserHandlers(userDatabase)

	v1 := app.Group("/api/v1")

	v1.Get("/users", userHandlers.GetUsers)
	v1.Post("/users/:name/entry", userHandlers.AddEntryToUser)
	v1.Post("/user", userHandlers.NewUser)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hellow World 👋 !")
	})

	initDatabase()

	setupRoutes(app)

	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(":3000"))
}
