package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/locator/dto"
	"github.com/guneyin/locator/mw"
	"github.com/guneyin/locator/service/location"
	"gorm.io/gorm"
)

const locationControllerName = "location"

type Location struct {
	svc *location.Service
}

func NewLocation(db *gorm.DB) IController {
	svc := location.New(db)

	return &Location{svc}
}

func (l Location) Name() string {
	return locationControllerName
}

func (l Location) SetRoutes(r fiber.Router) IController {
	g := r.Group(l.Name())
	g.Post("/", l.Add)
	g.Get("/", l.List)
	g.Get("/:id", l.Detail)
	g.Patch(":id", l.Edit)
	g.Post("/route", l.Route)
	g.Delete("/", l.Delete)

	return l
}

// Add
// @Summary Add new location.
// @Description Add new location to DB.
// @Tags location add
// @Accept json
// @Produce json
// @Param add body dto.LocationDto true "Add location"
// @Success 200 {object} mw.ResponseHTTP{data=dto.LocationResponseDto}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /location [post]
func (l Location) Add(c *fiber.Ctx) error {
	loc, err := dto.NewLocationDto(c.Body(), true)
	if err != nil {
		return err
	}

	res, err := l.svc.Add(c.Context(), loc)
	if err != nil {
		return err
	}

	return mw.OK(c, "location added successfully", res)
}

// List
// @Summary List locations.
// @Description List locations from DB.
// @Tags location list
// @Accept json
// @Produce json
// @Success 200 {object} mw.ResponseHTTP{data=dto.LocationListResponseDto}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /location [get]
func (l Location) List(c *fiber.Ctx) error {
	locList, err := l.svc.List(c.Context())
	if err != nil {
		return err
	}

	return mw.OK(c, "locations listed successfully", locList)
}

// Detail
// @Summary Location details.
// @Description Get location details from DB.
// @Tags location detail
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Success 200 {object} mw.ResponseHTTP{data=dto.LocationResponseDto}
// @Failure 404 {object} mw.ResponseHTTP{}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /location/{id} [get]
func (l Location) Detail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return mw.ErrInvalidId
	}

	loc, err := l.svc.Detail(c.Context(), uint(id))
	if err != nil {
		return err
	}

	return mw.OK(c, "location fetched successfully", loc)
}

// Edit
// @Summary Edit location.
// @Description Edit location data in DB.
// @Tags location edit
// @Accept json
// @Produce json
// @Param id path int true "Location ID"
// @Param location body dto.LocationDto true "Edit location"
// @Success 200 {object} mw.ResponseHTTP{data=dto.LocationResponseDto}
// @Failure 404 {object} mw.ResponseHTTP{}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /location/{id} [patch]
func (l Location) Edit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return mw.ErrInvalidId
	}

	body, err := dto.NewLocationDto(c.Body(), false)
	if err != nil {
		return err
	}

	loc, err := l.svc.Edit(c.Context(), uint(id), body)
	if err != nil {
		return err
	}

	return mw.OK(c, "location fetched successfully", loc)
}

// Route
// @Summary Route locations.
// @Description Route locations by given location.
// @Tags location route
// @Accept json
// @Produce json
// @Param location body dto.LocationDto true "Edit location"
// @Success 200 {object} mw.ResponseHTTP{data=dto.LocationListResponseDto}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /location/route [post]
func (l Location) Route(c *fiber.Ctx) error {
	loc, err := dto.NewLocationDto(c.Body(), true)
	if err != nil {
		return err
	}

	route, err := l.svc.Route(c.Context(), loc)
	if err != nil {
		return err
	}

	return mw.OK(c, "route listed successfully", route)
}

// Delete
// @Summary Delete locations.
// @Description Delete location data in DB.
// @Tags location delete
// @Accept json
// @Produce json
// @Param location body dto.LocationIdDto true "Delete locations"
// @Success 200 {object} mw.ResponseHTTP{}
// @Failure 404 {object} mw.ResponseHTTP{}
// @Failure 500 {object} mw.ResponseHTTP{}
// @Router /location [delete]
func (l Location) Delete(c *fiber.Ctx) error {
	idList := dto.LocationIdDto{}

	if err := c.BodyParser(&idList); err != nil {
		return err
	}

	err := l.svc.Delete(c.Context(), idList.IdList)
	if err != nil {
		return err
	}

	return mw.OK(c, "locations deleted successfully", nil)
}
