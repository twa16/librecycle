package main

import (
	"fmt"
	geo "github.com/kellydunn/golang-geo"
	"github.com/tkrajina/gpxgo/gpx"
	"github.com/twa16/librecycle/app/utils"
	"io/ioutil"
)

func main() {
	gpxBytes, err := ioutil.ReadFile("/home/mgauto/Downloads/Home-to-Mosaic.gpx")
	if err != nil {
		fmt.Println(err.Error())
	}
	gpxFile, err := gpx.ParseBytes(gpxBytes)

	stravaPoints := gpxFile.Tracks[0].Segments[0].Points

	geoPoints := utils.GPXPointsToGeoPoints(stravaPoints)

	for i, _ := range geoPoints {
		p1 := geoPoints[i]
		p2 := geoPoints[i+1]
		p3 := geoPoints[i+2]
		bearing := p1.BearingTo(&p2)

		fmt.Printf("%.5f:%.5f -> %.5f:%.5f == %.2f (%.9f)\n", p1.Lat(), p1.Lng(), p2.Lat(), p2.Lng(), bearing, dirOfPoint(p1, p2, p3))

	}
}

const (
	STRAIGHT = 0
	LEFT = -1
	RIGHT = 1
)
func dirOfPoint(a, b, p geo.Point) float64 {
	// subtracting co-ordinates of point A
	// from B and P, to make A as origin
	bLng := b.Lng() - a.Lng()
	bLat := b.Lat() - a.Lat()
	pLng := p.Lng() - a.Lng()
	pLat := p.Lat() - a.Lat()

	// Determining cross Product
	cross_product := bLng * pLat - bLat * pLng

	// return RIGHT if cross product is positive
	if (cross_product > 0) {
		return cross_product
	}

	// return LEFT if cross product is negative
	if (cross_product < 0){
		return cross_product
	}

	// return ZERO if cross product is zero.
	return cross_product
}