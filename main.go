package main

import "github.com/twa16/librecycle/app/bike/hardware/ebm"

func main() {
	bm := ebm.BikeModel{}
	bm.ScanAndConnect("00:1E:C0:7D:12:D0")
}
