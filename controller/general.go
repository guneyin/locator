package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/locator/controller/dto"
	"github.com/guneyin/locator/mw"
	"github.com/guneyin/locator/service/general"
	"log/slog"
)

const generalHandlerName = "general"

type GeneralHandler struct {
	svc *general.Service
}

var _ IController = (*GeneralHandler)(nil)

func NewGeneral(_ *slog.Logger) IController {
	svc := general.New()

	return &GeneralHandler{svc}
}

func (h GeneralHandler) Name() string {
	return generalHandlerName
}

func (h GeneralHandler) SetRoutes(r fiber.Router) IController {
	g := r.Group(h.Name())
	g.Get("status", h.GeneralStatus)

	return h
}

// Status
// @Summary Show the status of server.
// @Description Get the status of server.
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} middleware.ResponseHTTP{data=general.Status}
// @Failure 404 {object} middleware.ResponseHTTP{}
// @Failure 500 {object} middleware.ResponseHTTP{}
// @Router /general/status [get]
func (h GeneralHandler) GeneralStatus(c *fiber.Ctx) error {
	status := dto.StatusFromEntity(h.svc.Status())

	return mw.OK(c, "service status fetched", status)
}
