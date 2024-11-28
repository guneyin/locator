package dto

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidLocation    = errors.New("invalid location")
	ErrInvalidMarkerColor = errors.New("invalid marker color")
)

type LocationDto struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Name        string  `json:"name"`
	MarkerColor string  `json:"markerColor"`
}

type LocationResponseDto struct {
	Id string `json:"id"`
	LocationDto
}

type LocationListResponseDto []LocationResponseDto

func NewLocationDto(data []byte) (*LocationDto, error) {
	loc := &LocationDto{}

	err := json.Unmarshal(data, loc)
	if err != nil {
		return nil, err
	}

	return loc.validate()
}

func (l *LocationDto) validate() (*LocationDto, error) {
	var errs error

	if l.Latitude <= 0 || l.Longitude <= 0 {
		errs = errors.Join(errs, ErrInvalidLocation)
	}
	if l.MarkerColor == "" {
		errs = errors.Join(errs, ErrInvalidMarkerColor)
	}

	return l, errs
}
