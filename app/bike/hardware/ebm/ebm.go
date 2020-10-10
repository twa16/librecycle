package ebm

import (
	"fmt"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/twa16/librecycle/app/bike/hardware/ebm/ebmproto"
)

type BikeModel struct {
	MACAddress string
	ShouldRun bool
	BatteryPercent int
	gattDevice *gatt.Device
	LatestBatteryMessage ebmproto.BatteryMessage
	LatestMotorMessage ebmproto.MotorMessage
	OnStateUpdate func(model *BikeModel)
}

func (bm *BikeModel) ScanAndConnect(bikeMAC string) error {
	bm.MACAddress = bikeMAC

	dev, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		return err
	}

	// Register handlers.
	dev.Handle(gatt.PeripheralDiscovered(bm.onPeriphDiscovered),
		gatt.PeripheralConnected(bm.onPeriphConnected),
		gatt.PeripheralDisconnected(bm.onPeriphDisconnected))
	dev.Init(bm.onStateChanged)

	bm.ShouldRun = true
	for bm.ShouldRun {}
	return nil
}

func (bm *BikeModel) onStateChanged(d gatt.Device, s gatt.State) {
	switch s {
	case gatt.StatePoweredOn:
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

func (bm *BikeModel) onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	if p.ID() == bm.MACAddress {
		fmt.Println("Found Bike! Connecting!")
		p.Device().Connect(p)
	} else {
		fmt.Printf("Found Device %s (%s)\n", p.ID(), p.Name())
	}
}

func (bm *BikeModel) onPeriphConnected(p gatt.Peripheral, err error) {
	if err := p.SetMTU(500); err != nil {
		fmt.Printf("Failed to set MTU, err: %s\n", err)
	}

	// Discovery services
	ss, err := p.DiscoverServices(nil)
	if err != nil {
		fmt.Printf("Failed to discover services, err: %s\n", err)
		return
	}

	for _, s := range ss {
		msg := "Service: " + s.UUID().String()
		if len(s.Name()) > 0 {
			msg += " (" + s.Name() + ")"
		}
		fmt.Println(msg)

		// Discovery characteristics
		cs, err := p.DiscoverCharacteristics(nil, s)
		if err != nil {
			fmt.Printf("Failed to discover characteristics, err: %s\n", err)
			continue
		}

		for _, c := range cs {
			msg := "  Characteristic  " + c.UUID().String()
			if len(c.Name()) > 0 {
				msg += " (" + c.Name() + ")"
			}
			msg += "\n    properties    " + c.Properties().String()
			fmt.Println(msg)

			// Read the characteristic, if possible.
			if (c.Properties() & gatt.CharRead) != 0 {
				b, err := p.ReadCharacteristic(c)
				if err != nil {
					fmt.Printf("Failed to read characteristic, err: %s\n", err)
					continue
				}
				fmt.Printf("    value         %x | %q\n", b, b)
			}

			// Discovery descriptors
			ds, err := p.DiscoverDescriptors(nil, c)
			if err != nil {
				fmt.Printf("Failed to discover descriptors, err: %s\n", err)
				continue
			}

			for _, d := range ds {
				msg := "  Descriptor      " + d.UUID().String()
				if len(d.Name()) > 0 {
					msg += " (" + d.Name() + ")"
				}
				fmt.Println(msg)

				// Read descriptor (could fail, if it's not readable)
				b, err := p.ReadDescriptor(d)
				if err != nil {
					fmt.Printf("Failed to read descriptor, err: %s\n", err)
					continue
				}
				fmt.Printf("    value         %x | %q\n", b, b)
			}

			// Subscribe the characteristic, if possible.
			if (c.Properties() & (gatt.CharNotify | gatt.CharIndicate)) != 0 {
				f := func(c *gatt.Characteristic, b []byte, err error) {
					//fmt.Printf("notified: % X | %q\n", b, b)
					bm.onBLEMessage(b)
				}
				if err := p.SetNotifyValue(c, f); err != nil {
					fmt.Printf("Failed to subscribe characteristic, err: %s\n", err)
					continue
				}
			}

		}
		fmt.Println()
	}

	for bm.ShouldRun {}
}

func (bm *BikeModel) onPeriphDisconnected(p gatt.Peripheral, err error) {

}

func (bm *BikeModel) onBLEMessage(msgBytes []byte) {
	if msgBytes[1] == 0x62 {
		batMsg := ebmproto.BatteryMessage{}
		batMsg.ParseFromBytes(msgBytes)
		bm.LatestBatteryMessage = batMsg
	} else if msgBytes[1] == 0x6d && msgBytes[3] == 0x5a {
		motMsg := ebmproto.MotorMessage{}
		motMsg.ParseFromBytes(msgBytes)
		bm.LatestMotorMessage = motMsg
	} else {
		fmt.Printf("Got Msg of Types: %x %x\n", msgBytes[1], msgBytes[3])
	}
}

func (bm *BikeModel) onStateUpdate() {
	bm.OnStateUpdate(bm)
}