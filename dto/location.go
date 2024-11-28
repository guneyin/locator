package dto

import (
	"encoding/json"
	"errors"
	"github.com/guneyin/locator/repository/location"
	"github.com/guneyin/locator/util"
)

var (
	ErrInvalidLocation    = errors.New("invalid location")
	ErrInvalidMarkerColor = errors.New("invalid marker color")
)

type LocationDto struct {
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Name        string `json:"name"`
	MarkerColor string `json:"markerColor"`
}

type LocationResponseDto struct {
	Id uint `json:"id"`
	LocationDto
}

type LocationListResponseDto struct {
	Items []LocationResponseDto
}

func NewLocationDto(data []byte) (*LocationDto, error) {
	loc := &LocationDto{}

	err := json.Unmarshal(data, loc)
	if err != nil {
		return nil, err
	}

	return loc.validate()
}

func NewLocationResponseDto(entity *location.Location) (*LocationResponseDto, error) {
	loc := &LocationResponseDto{
		Id: entity.ID,
		LocationDto: LocationDto{
			Latitude:    util.FloatToStr(entity.Latitude),
			Longitude:   util.FloatToStr(entity.Longitude),
			Name:        entity.Name,
			MarkerColor: entity.MarkerColor,
		},
	}

	return loc, nil
}

func NewLocationListResponseDto(entity location.LocationList) (*LocationListResponseDto, error) {
	locList := make([]LocationResponseDto, len(entity))

	for i, item := range entity {
		locList[i] = LocationResponseDto{
			Id: item.ID,
			LocationDto: LocationDto{
				Latitude:    util.FloatToStr(item.Latitude),
				Longitude:   util.FloatToStr(item.Longitude),
				Name:        item.Name,
				MarkerColor: item.MarkerColor,
			},
		}
	}

	return &LocationListResponseDto{Items: locList}, nil
}

func (l *LocationDto) validate() (*LocationDto, error) {
	var errs error

	if l.Latitude == "0" || l.Longitude == "0" {
		errs = errors.Join(errs, ErrInvalidLocation)
	}
	if l.MarkerColor == "" {
		errs = errors.Join(errs, ErrInvalidMarkerColor)
	}

	return l, errs
}

func (l *LocationDto) ToEntity() *location.Location {
	lat, _ := util.StrToFloat(l.Latitude)
	lon, _ := util.StrToFloat(l.Longitude)

	return &location.Location{
		Latitude:    lat,
		Longitude:   lon,
		Name:        l.Name,
		MarkerColor: l.MarkerColor,
	}
}