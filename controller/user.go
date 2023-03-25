package controller


import (
	"strings"
	"time"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/MBI-88/fiberApp/model"
	"github.com/golang-jwt/jwt/v5"
)


var keypassword = []byte("use123#!na87)(12|°)")

// UserController struct
type UserController struct{}

// ConfigPath url to get user
func (u *UserController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/",u.ValidateToken,u.GetUser)
	app.Post("/", u.RegisterUser)
	app.Post("/login",u.HandlerLogin)
	return app
}

// GetUser method for getting users
func (u *UserController) GetUser(c *fiber.Ctx) error {
	modelUser := model.User{}
	users, err := modelUser.GetUsers()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	return c.JSON(users)

}

// RegisterUser method
func (u *UserController) RegisterUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	if err := user.RegisterUser(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	return c.Status(http.StatusAccepted).JSON("Accepted!")
}

// HandlerLogin user
func (u *UserController) HandlerLogin(c *fiber.Ctx) error {
	var user model.User 
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	err := user.Login()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	tokenString,err := GenerateToken(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}

	response := model.JWTClient{
		Name: user.Name,
		Token: tokenString,
	}
	return c.Status(http.StatusAccepted).JSON(response)
}

// ValidateToken method
func (u *UserController) ValidateToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization","")
	token := strings.Split(tokenString,"Bearer")
	if len(token) != 2 {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: "Error token"})
	}
	tokenCleaned := strings.TrimSpace(token[1])
	claims := model.Claims{}
	tokenParsed, err := jwt.ParseWithClaims(tokenCleaned,&claims,
		func(t *jwt.Token) (interface{},error) {
			return keypassword,nil})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	if !tokenParsed.Valid{
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: "Token no valid"})
	}
	return c.Next()
}

// GenerateToken function
func GenerateToken(user model.User) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id": user.UserID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	})
	return token.SignedString(keypassword)
}

// NewUserController return users
func NewUserController() *UserController {
	return &UserController{}
}
