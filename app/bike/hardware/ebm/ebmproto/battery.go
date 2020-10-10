package ebmproto

import (
	"encoding/binary"
)

type BatteryMessage struct {
	Current float64
	CurrentCapacity float64
	Cycles int
	NominalCapacity float64
	Percent int
	Temperature int //Celsius
	Voltage float64
}

func (m *BatteryMessage) ParseFromBytes(bytes []byte) {
	m.Voltage = convBytesToFloat(bytes[5], bytes[6])
	m.Percent = int(bytes[7])
	m.Temperature = int(bytes[8])
	m.Current = convBytesToFloat(bytes[9], bytes[10])
	m.NominalCapacity = convBytesToFloat(bytes[11], bytes[12])
	m.CurrentCapacity = convBytesToFloat(bytes[13], bytes[14])
	if len(bytes) >= 19 {
		//intermediate := convBytesToFloat(bytes[15], bytes[16])
		//m.Cycles = int(intermediate)
		m.Cycles = int(binary.BigEndian.Uint16(bytes[15:17])) - 10000
	}
}

