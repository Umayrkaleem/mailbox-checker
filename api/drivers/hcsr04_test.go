package hcsr04

import (
	"machine"
	"testing"
)

func TestSensor_GetDistance(t *testing.T) {
	type fields struct {
		trigger machine.Pin
		echo    machine.Pin
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sensor := &Sensor{
				trigger: tt.fields.trigger,
				echo:    tt.fields.echo,
			}
			if got := Sensor.GetDistance(); got != tt.want {
				t.Errorf("Sensor.GetDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
