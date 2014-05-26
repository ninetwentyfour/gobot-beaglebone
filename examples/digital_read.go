package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-beaglebone"
	"github.com/hybridgroup/gobot-gpio"
)

func main() {
	beaglebone := new(gobotBeaglebone.Beaglebone)
	beaglebone.Name = "beaglebone"

	button := gobotGPIO.NewDirectPin(beaglebone)
	button.Name = "button"
	button.Pin = "P8_9"

	work := func() {
		gobot.Every("1s", func() {
			fmt.Println(button.DigitalRead())
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{beaglebone},
		Devices:     []gobot.Device{button},
		Work:        work,
	}

	robot.Start()
}
