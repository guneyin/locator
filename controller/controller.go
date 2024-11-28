package controller

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type IController interface {
	Name() string
	SetRoutes(r fiber.Router) IController
}

type Controller struct {
	log         *slog.Logger
	router      fiber.Router
	controllers map[string]IController
}

func New(log *slog.Logger, router fiber.Router) *Controller {
	handler := &Controller{
		log:         log,
		router:      router,
		controllers: make(map[string]IController),
	}
	handler.registerHandlers()

	return handler
}

func (c Controller) registerHandlers() {
	c.registerHandler(NewGeneral)
}

func (c Controller) registerHandler(f func(log *slog.Logger) IController) {
	hnd := f(c.log).SetRoutes(c.router)
	c.controllers[hnd.Name()] = hnd
}
