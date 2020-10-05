package main

import (
	"context"
	"fmt"
	"github.com/tkrajina/gpxgo/gpx"
	"github.com/twa16/librecycle/app/utils"
	"googlemaps.github.io/maps"
	"io/ioutil"
	"github.com/kr/pretty"
	"log"
	"os"
)

func main() {
	pathToGpx := "/Users/mgauto/Downloads/Home-to-Mosaic.gpx"
	gpxBytes, err := ioutil.ReadFile(pathToGpx)
	if err != nil {
		fmt.Println(err.Error())
	}
	gpxFile, err := gpx.ParseBytes(gpxBytes)

	stravaPoints := gpxFile.Tracks[0].Segments[0].Points

	geoPoints := utils.GPXPointsToGeoPoints(stravaPoints)


	for i, _ := range geoPoints[:len(geoPoints)-2] {
		p1 := geoPoints[i]
		p2 := geoPoints[i+1]
		p3 := geoPoints[i+2]
		bearing := p2.BearingTo(&p3)
		adjustedBearing := bearing - p1.BearingTo(&p2)

		if p2.GreatCircleDistance(&p3) < .05 {
			fmt.Printf("Points Close. Ignoring. Dist: %.3f\n", p2.GreatCircleDistance(&p3))
		} else {
			//fmt.Printf("%.5f:%.5f -> %.5f:%.5f == %.2f (%.9f)\n", p1.Lat(), p1.Lng(), p2.Lat(), p2.Lng(), bearing, dirOfPoint(p1, p2, p3))
			if (adjustedBearing < 5 && adjustedBearing > 0) || (adjustedBearing > -5 && adjustedBearing <= 0)  {
				fmt.Printf("Bearing is small. Ignoring. Bearing: %.3f\n", adjustedBearing)
			} else {
				fmt.Printf("TURN %.2f\n", adjustedBearing)
			}
		}
	}

	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := maps.NearestRoadsRequest{Points: utils.GeoPointsToGmapLatLngs(geoPoints)}

	resp, err := c.NearestRoads(context.Background(), &r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(resp.SnappedPoints)
}
