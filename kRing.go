package H3Go

import (
	"math"

	"github.com/uber/h3-go"
)

/*
	H3Index returns the index of the cell of the location
	PARAMS latitude, longitude float64, res - H3 cell resolution
*/
func H3Index(latitude, longitude float64, res int) h3.H3Index {
	coord := h3.GeoCoord{
		Latitude:  latitude,
		Longitude: longitude,
	}
	index := h3.FromGeo(coord, res)
	return index
}

/*
	KRing returns the array of neighbouring cell indexes at a given distance KM
	PARAMS origin base cell index, resolution int, distance in KM
*/
func KRing(origin h3.H3Index, res int, distance float64) []h3.H3Index {
	radius := math.Floor(distance / ((cellEdgeLength[res]) * 2))
	neighbours := h3.KRing(origin, int(radius))
	return neighbours
}
