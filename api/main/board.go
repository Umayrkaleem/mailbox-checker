package main

import (
	"github.com/Umayrkaleem/mailbox-checker/api/drivers"
	"machine"
	"time"
)

func main() {
	sensor := hcsr04.New(machine.Pin(2), machine.Pin(3))
	sensor.Configure()

	for {
		println(sensor.GetDistance())
		time.Sleep(time.Second)
	}

}
