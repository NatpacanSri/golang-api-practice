package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct{
	ID 			int 	`json:"id"`
	FirstName 	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
}

var users = []User{
	{ID:1, FirstName:"John", LastName:"Doe"},
	{ID:2, FirstName:"Jame", LastName:"Dan"},
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error{
	id := c.Params("id")
	for _, user := range users {
		if strconv.Itoa(user.ID) == id {
			return c.JSON(user)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Invalid"})
	}
	user.ID = len(users) +1
	users = append(users, *user)
	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			updateUser := new(User)
			if err := c.BodyParser(updateUser); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid"})
			}
			updateUser.ID = user.ID
			users[i] = *updateUser
			return c.JSON(updateUser)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
}

func main() {
	
	app := fiber.New()

	app.Get("/users",getUsers)
	app.Get("/user",getUser)
	app.Get("/user/:id",getUser)
	app.Post("/createUser",createUser)
	app.Put("/updateUser/:id",updateUser)
	app.Delete("/deleteUser/:id",deleteUser)

	err := app.Listen(":8000")
	if err != nil{
		panic(err)
	}
}