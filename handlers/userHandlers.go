package handlers

import (
	"fmt"
	"strconv"

	"github.com/j1018813/doeless/models/user"

	"github.com/gofiber/fiber/v2"
	"github.com/j1018813/doeless/database"
)

// UserHandlers the handlers for the User model
type UserHandlers struct {
	userDb *database.UserDatabase
}

// NewUserHandlers gives you a User handler
func NewUserHandlers(uDb *database.UserDatabase) *UserHandlers {
	return &UserHandlers{
		userDb: uDb,
	}
}

// GetUsers is the handler method which retrieves all users
func (u *UserHandlers) GetUsers(c *fiber.Ctx) error {
	users, err := u.userDb.GetUsers()
	c.JSON(users)
	return err
}

// GetUser is the handler method which retrieves user with given id
func (u *UserHandlers) GetUser(c *fiber.Ctx) error {
	users, err := u.userDb.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		// string to int
		i, err := strconv.ParseUint(c.Params("id"), 0, 16)
		if err != nil {
			// handle error
			fmt.Println(err)
			return err
		}

		if user.ID == uint(i) {
			c.JSON(user)
			break
		}
	}

	return err
}

// NewUser is the handler method which creates a new user
func (u *UserHandlers) NewUser(c *fiber.Ctx) error {
	user := new(user.User)

	if err := c.BodyParser(user); err != nil {
		c.Status(503).SendString("failed to parse body to user")
		return err
	}

	u.userDb.NewUser(user)
	c.JSON(user)
	return nil
}

// AddEntryToUser adds an entry to the given user
func (u *UserHandlers) AddEntryToUser(c *fiber.Ctx) error {
	fmt.Fprintf(c, "%s\n", c.Params("name"))
	return nil
}
