package config

import (
	"github.com/gofiber/fiber/v2"
)

// Microservice interface for config
// microservices
type Microservice interface {
	ConfigPath(app *fiber.App) *fiber.App 

}

// User microservice
func Use(prefix string, r fiber.Router,micro Microservice) {
	r.Mount(prefix,micro.ConfigPath(fiber.New()))
}