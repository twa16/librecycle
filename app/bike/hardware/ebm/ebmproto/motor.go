package ebmproto

type MotorMessage struct {
	AssistLevel int
	Cadence int
	Current float64
	MaxCurrent float64
	MaxTorque int
	Speed float64
	Temperature int
	Torque int
}

func (mm *MotorMessage) ParseFromBytes(b []byte) {
	mm.AssistLevel = int(b[5])
	mm.Temperature = int(b[6])
	mm.Current = convBytesToFloat(b[8], b[7]) / 10.0
	mm.Speed = convBytesToFloat(b[10], b[9]) / 10.0
	mm.Cadence = int(b[11])
	mm.Torque = int(b[12])
	mm.MaxCurrent = convBytesToFloat(b[13], b[14]) / 10.0
	mm.MaxTorque = int(b[15])
}