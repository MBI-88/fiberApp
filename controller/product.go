package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/MBI-88/fiberApp/model"
	"net/http"
)

// ProductController struct
type ProductController struct{}

// ConfigPath method router
func (p *ProductController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/",p.HandleGetProduct)
	app.Post("/",p.HandleSaveProduct)
	return app
}

// HandleGetProduct menthod for getting products
func (p *ProductController) HandleGetProduct(c *fiber.Ctx) error {
	product := model.Product{}
	products, err := product.GetProducts()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.JSON(products)
}

// HandleSaveProduct method for saving products
func (p *ProductController) HandleSaveProduct(c *fiber.Ctx) error {
	var product model.Product 
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	if err := product.SaveProducts(); err != nil {
		return c.Status(http.StatusConflict).JSON(model.Errors{
			Message: err.Error()})
	}

	return c.Status(200).JSON(product)
}

// NewProductController function
// generate ProductController
func NewProductController() * ProductController {
	return &ProductController{}
}