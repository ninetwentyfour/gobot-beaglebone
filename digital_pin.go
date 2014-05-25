package gobotBeaglebone

import (
	"os"
	"strconv"
	"strings"
)

type digitalPin struct {
	PinNum  string
	Mode    string
	PinFile *os.File
	Status  string
}

const GPIO_PATH = "/sys/class/gpio"
const GPIO_DIRECTION_READ = "in"
const GPIO_DIRECTION_WRITE = "out"
const HIGH = 1
const LOW = 0

func newDigitalPin(pinNum int, mode string) *digitalPin {
	d := new(digitalPin)
	d.PinNum = strconv.Itoa(pinNum)

	fi, err := os.OpenFile(GPIO_PATH+"/export", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	fi.WriteString(d.PinNum)
	fi.Close()

	d.setMode(mode)

	return d
}

func (d *digitalPin) setMode(mode string) {
	d.Mode = mode

	if mode == "w" {
		fi, err := os.OpenFile(GPIO_PATH+"/gpio"+d.PinNum+"/direction", os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		fi.WriteString(GPIO_DIRECTION_WRITE)
		fi.Close()
		d.PinFile, err = os.OpenFile(GPIO_PATH+"/gpio"+d.PinNum+"/value", os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
	} else if mode == "r" {
		fi, err := os.OpenFile(GPIO_PATH+"/gpio"+d.PinNum+"/direction", os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		fi.WriteString(GPIO_DIRECTION_READ)
		fi.Close()
	}
}

func (d *digitalPin) digitalWrite(value string) {
	if d.Mode != "w" {
		d.setMode("w")
	}

	d.PinFile.WriteString(value)
	d.PinFile.Sync()
}

func (d *digitalPin) digitalRead() int {
	if d.Mode != "r" {
		d.setMode("r")
	}

	var err error
	d.PinFile, err = os.OpenFile(GPIO_PATH+"/gpio"+d.PinNum+"/value", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	var buf []byte = make([]byte, 1024)
	d.PinFile.Read(buf)
	d.PinFile.Close()

	i, _ := strconv.Atoi(strings.Split(string(buf), "\n")[0])
	return i
}

func (d *digitalPin) close() {
	fi, err := os.OpenFile(GPIO_PATH+"/unexport", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	fi.WriteString(d.PinNum)
	fi.Close()
	d.PinFile.Close()
}
