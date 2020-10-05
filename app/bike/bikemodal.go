package bike

import (
	geo "github.com/kellydunn/golang-geo"
	"time"
)

type Model struct {
	Name string
	CurLocation geo.Point
	LastLocationUpdate time.Time
	Speed float64
	Heading float64
}
