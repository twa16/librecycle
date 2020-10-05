package turnbyturn

import (
	"fmt"
	geo "github.com/kellydunn/golang-geo"
)

const (
	POINT_EPSILON = 0.05 //Distance between points to trigger look ahead check
	BEARING_EPSILON = 5 //Size of turn to ignore (Positive Degrees)
)

type Maneuver struct {
	Origin geo.Point
	NewBearing float64
}

func FindFirstTurn([]geo.Point) {

}

func FindNextTurn(curPos geo.Point, route[]geo.Point) {
	for i, _ := range route[:len(route)-1] {
		p2 := route[i]
		p3 := route[i+1]
		bearing := p2.BearingTo(&p3)
		adjustedBearing := bearing - curPos.BearingTo(&p2)

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
}