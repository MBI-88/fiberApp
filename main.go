package main


import (
	"github.com/MBI-88/fiberApp/controller"
	"github.com/MBI-88/fiberApp/config"
	"github.com/gofiber/fiber/v2"
	"github.com/MBI-88/fiberApp/model"
)



func main() {
	app := fiber.New()
	model.InitDB()

	//ctrUser := controller.NewUserController()
	//app = ctrUser.ConfigPath(app)

	apiV1 := app.Group("/apiV1")
	config.Use("/user",apiV1,controller.NewUserController()) // apiV1/user/
	app.Listen(":8000")

	
}