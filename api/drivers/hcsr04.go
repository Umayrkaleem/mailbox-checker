package hcsr04

import (
	"machine"
	"time"
)

type Sensor struct {
	trigger machine.Pin
	echo    machine.Pin
}

func New(trigger, echo machine.Pin) Sensor {
	return Sensor{
		trigger: trigger,
		echo:    echo,
	}
}

func (sensor *Sensor) Configure() {
	sensor.trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
}

func (Sensor *Sensor) GetDistance() int32 {

	// Trigger pin sends a pulse
	Sensor.trigger.Low()
	time.Sleep(time.Microsecond * 2)
	Sensor.trigger.High()
	time.Sleep(time.Microsecond * 10)
	Sensor.trigger.Low()

	// Measure the duration it takes to receive pulse on the echo pin
	start := time.Now()
	for Sensor.echo.Get() == true {
		if time.Since(start) > 100*time.Millisecond {
			return 0
		}
	}
	end := time.Now()

	for Sensor.echo.Get() == true {
		if time.Since(end) > 100*time.Millisecond {
			return 1
		}
	}

	duration := end.Sub(start)
	distance := int32(duration.Seconds() * 34300 / 2)

	return int32(distance)
}
