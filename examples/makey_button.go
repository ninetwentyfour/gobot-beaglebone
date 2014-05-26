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

	button := gobotGPIO.NewMakeyButton(beaglebone)
	button.Name = "button"
	button.Pin = "P8_9"

	work := func() {
		gobot.On(button.Events["push"], func(data interface{}) {
			fmt.Println("button pressed")
		})

		gobot.On(button.Events["release"], func(data interface{}) {
			fmt.Println("button released")
		})
	}

	robot := gobot.Robot{
		Connections: []gobot.Connection{beaglebone},
		Devices:     []gobot.Device{button},
		Work:        work,
	}

	robot.Start()
}
