package controller

import (
	"net/http"
	"strconv"
	"github.com/MBI-88/fiberApp/model"
	"github.com/gofiber/fiber/v2"
)

// FactureController struct
type FactureController struct{}

// ConfigPath method router
func (f *FactureController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/",f.HandleSaveFacture)
	app.Post("/",f.HandleSaveFacture)
	app.Get("/:id",f.HandleSaveFacture)
	app.Get("/user/:id",f.HandleFactureUser)
	return app
}

// HandleSaveFacture method for saving facture
func (f *FactureController) HandleSaveFacture(c *fiber.Ctx) error {
	var facture model.Facture
	if err := c.BodyParser(&facture); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	if err := facture.SaveFacture(); err != nil {
		return c.Status(http.StatusConflict).JSON(model.Errors{
			Message: err.Error()})
	}

	return c.Status(200).JSON("Accepted!")
}

// HandleGetFactures method to return all factures
func (f FactureController) HandleGetFactures(c *fiber.Ctx) error {
	factures,err := model.Facture{}.GetFactures()
	if err != nil {
		return c.Status(http.StatusConflict).JSON(model.Errors{
			Message: err.Error()})
	}
	return c.Status(200).JSON(factures)
}

// HandleGetFactureID method to get facture by id
func (FactureController) HandleGetFactureID(c *fiber.Ctx) error {
	idstring := c.Params("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	var facture model.Facture
	er :=  facture.GetFacture(uint(id))
	if er != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: er.Error()})
	}
	return c.Status(200).JSON(facture)
}

// HandleFactureUser menthod to return facture by user
func (FactureController) HandleFactureUser(c *fiber.Ctx) error {
	idstring := c.Params("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}
	factures,err := model.Facture{}.GetFactureUser(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Errors{
			Message: err.Error()})
	}

	return c.Status(200).JSON(factures)
}

// NewFactureController function generate FactureController
func NewFactureController() *FactureController {
	return &FactureController{}
}

