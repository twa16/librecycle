package datalogger

import (
	"github.com/twa16/librecycle/app/bike"
	"time"
)

type DataPoint struct {
	Timestamp          time.Time
	BikeModelSnapshot  bike.BikeModel
	VectorDataPoints   []VectorDataPoint
	LocationDataPoints []LocationDataPoint
	AdditionalValues   map[string]string
}

type LocationDataPoint struct {
	Source         string
	Latitude       float64
	Longitude      float64
	RawSource      []byte
	AdditionalData map[string]string
}

type VectorDataPoint struct {
	Source    string
	Speed     float64
	Heading   float64
	RawSource []byte
}
