package utils

import (
	"github.com/kellydunn/golang-geo"
	"github.com/tkrajina/gpxgo/gpx"
	"googlemaps.github.io/maps"
)

func GPXPointsToGeoPoints(gpxPoints []gpx.GPXPoint) []geo.Point {
	result := make([]geo.Point, len(gpxPoints))
	for i, origPoint := range gpxPoints {
		result[i] = *geo.NewPoint(origPoint.Latitude, origPoint.Longitude)
	}
	return result
}

func GeoPointsToGmapLatLngs(points []geo.Point) []maps.LatLng {
	result := make([]maps.LatLng, len(points))
	for i, origPoint := range points {
		result[i] = maps.LatLng{Lat: origPoint.Lat(), Lng: origPoint.Lng()}
	}
	return result
}