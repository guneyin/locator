package location

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Name        string  `json:"name"`
	MarkerColor string  `json:"markerColor"`
}

type LocationList []Location
