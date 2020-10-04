package utils

import (
	"github.com/kellydunn/golang-geo"
	"github.com/tkrajina/gpxgo/gpx"
)

func GPXPointsToGeoPoints(gpxPoints []gpx.GPXPoint) []geo.Point {
	result := make([]geo.Point, len(gpxPoints))
	for i, origPoint := range gpxPoints {
		result[i] = *geo.NewPoint(origPoint.Latitude, origPoint.Longitude)
	}
	return result
}
