package controller


import (
	"github.com/gofiber/fiber/v2"
	"github.com/MBI-88/fiberApp/model"
)


// UserController struct
type UserController struct{}

// ConfigPath url to get user
func (u *UserController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/",u.GetUser)
	return app
}

// GetUser method for getting users
func (u *UserController) GetUser(c *fiber.Ctx) error {
	modelUser := model.User{}
	users, err := modelUser.GetUsers()
	if err != nil {
		return c.SendString("Error")
	}
	return c.JSON(users)

}

// NewUserController return users
func NewUserController() *UserController {
	return &UserController{}
}
