package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-beaglebone"
	"github.com/hybridgroup/gobot-gpio"
	"time"
)

func main() {
	beaglebone := new(gobotBeaglebone.Beaglebone)
	beaglebone.Name = "beaglebone"

	led := gobotGPIO.NewDirectPin(beaglebone)
	led.Name = "led"
	led.Pin = "P8_10"

	work := func() {
		for {
			led.DigitalWrite(1)
			time.Sleep(1000 * time.Millisecond)
			led.DigitalWrite(0)
			time.Sleep(1000 * time.Millisecond)
		}
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{beaglebone},
		Devices:     []gobot.Device{led},
		Work:        work,
	}

	robot.Start()
}
