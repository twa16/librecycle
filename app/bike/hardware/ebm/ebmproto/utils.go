package ebmproto

import "math"

func convBytesToFloat(b1R byte, b2R byte) float64 {
	b1 := uint16(b1R)
	b2 := uint16(b2R)

	noDecimal := float64( b1 << 8 | b2 )

	return math.Floor( (noDecimal/ 10.0)*100)/100
}
