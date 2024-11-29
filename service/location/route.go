package location

import (
	"github.com/guneyin/locator/dto"
	"github.com/guneyin/locator/repository/location"
	"github.com/guneyin/locator/util"
	"sort"
)

type Route struct {
	Items RouteItemList
}

type RouteItemList []RouteItem

type RouteItem struct {
	Id          int8
	Latitude    float64
	Longitude   float64
	Name        string
	MarkerColor string
	Distance    float64
	Order       int8
}

func (ri *RouteItem) IsMarked() bool {
	return ri.Distance > 0.0
}

func NewRoute(locList location.LocationList) *Route {
	items := make(RouteItemList, len(locList))
	for i, loc := range locList {
		items[i] = RouteItem{
			Id:          int8(loc.ID),
			Latitude:    loc.Latitude,
			Longitude:   loc.Longitude,
			Name:        loc.Name,
			MarkerColor: loc.MarkerColor,
			Distance:    0,
		}
	}

	return &Route{Items: items}
}

func (r *Route) Do(loc *dto.LocationDto) *Route {
	lat := util.StrToFloat(loc.Latitude)
	long := util.StrToFloat(loc.Longitude)
	shortestIndex := -1

	for i := range r.Items {
		r.Items[i].Distance = 0.0
	}

	for order := range len(r.Items) {
		shortestDistance := 10000.00
		for i, item := range r.Items {
			if item.IsMarked() {
				continue
			}

			distance := util.Haversine(lat, long, item.Latitude, item.Longitude)
			if distance < shortestDistance {
				shortestIndex = i
				shortestDistance = distance
			}
		}

		r.Items[shortestIndex].Distance = shortestDistance
		r.Items[shortestIndex].Order = int8(order)

		lat = r.Items[shortestIndex].Latitude
		long = r.Items[shortestIndex].Longitude
	}

	routeItemList := make([]RouteItem, 1, len(r.Items)+1)
	routeItemList[0] = RouteItem{
		Id:          -1,
		Latitude:    util.StrToFloat(loc.Latitude),
		Longitude:   util.StrToFloat(loc.Longitude),
		Name:        loc.Name,
		MarkerColor: loc.MarkerColor,
		Distance:    0,
		Order:       -1,
	}
	r.Items = append(routeItemList, r.Items...)

	return r
}

func (r *Route) ToLocationListResponseDto() *dto.LocationListResponseDto {
	locList := &dto.LocationListResponseDto{
		Items: make([]dto.LocationResponseDto, len(r.Items)),
	}

	sort.Slice(r.Items, func(i, j int) bool {
		return r.Items[i].Order < r.Items[j].Order
	})

	for i, item := range r.Items {
		locList.Items[i] = dto.LocationResponseDto{
			Id: uint(item.Id),
			LocationDto: dto.LocationDto{
				Latitude:    util.FloatToStr(item.Latitude),
				Longitude:   util.FloatToStr(item.Longitude),
				Name:        item.Name,
				MarkerColor: item.MarkerColor,
			},
		}
	}

	return locList
}
