package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/locator/dto"
	"github.com/guneyin/locator/mw"
	"github.com/guneyin/locator/service/general"
	"gorm.io/gorm"
)

const generalControllerName = "general"

type General struct {
	svc *general.Service
}

func NewGeneral(_ *gorm.DB) IController {
	svc := general.New()

	return &General{svc}
}

func (g General) Name() string {
	return generalControllerName
}

func (g General) SetRoutes(r fiber.Router) IController {
	gr := r.Group(g.Name())
	gr.Get("status", g.GeneralStatus)

	return g
}

// Status
// @Summary Show the status of server.
// @Description Get the status of server.
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} mw.ResponseHTTP{data=general.Status}
// @Failure 404 {object} mw.ResponseHTTP{}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /general/status [get]
func (g General) GeneralStatus(c *fiber.Ctx) error {
	status := dto.StatusFromEntity(g.svc.Status())

	return mw.OK(c, "service status fetched", status)
}
