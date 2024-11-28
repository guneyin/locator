package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IController interface {
	Name() string
	SetRoutes(r fiber.Router) IController
}

type Controller struct {
	db          *gorm.DB
	router      fiber.Router
	controllers map[string]IController
}

func New(db *gorm.DB, router fiber.Router) *Controller {
	c := &Controller{
		db:          db,
		router:      router,
		controllers: make(map[string]IController),
	}
	c.registerControllers()

	return c
}

func (c Controller) registerControllers() {
	c.register(NewGeneral)
	c.register(NewLocation)
}

func (c Controller) register(f func(db *gorm.DB) IController) {
	hnd := f(c.db).SetRoutes(c.router)
	c.controllers[hnd.Name()] = hnd
}
