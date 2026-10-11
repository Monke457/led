package main

import (
	"time"

	"github.com/rpi-ws281x/rpi-ws281x-go"
)

func main() {
	cOpt := ws2811.ChannelOption{
		GpioPin : 18,
		Invert : false,
		LedCount : 256,
		Brightness : 8,
	}

	opt := ws2811.Option{
		Frequency:800000,
		Channels: []ws2811.ChannelOption{ cOpt },
	}

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}

	err = dev.Init()
	if err != nil {
		panic(err)
	}

	defer dev.Fini()

	setLed(dev, 0, 0, 0x00FF0000)
	renderLeds(dev)

	time.Sleep(time.Second * 3)

	setLed(dev, 0, 0, 0x00000000)
	renderLeds(dev)
}

func setLed(dev *ws2811.WS2811, channel int, led int, col uint32) {
	leds := dev.Leds(channel)
	leds[led] = col
}

func renderLeds(dev *ws2811.WS2811) { 
	if err := dev.Render(); err != nil {
		panic(err)
	}
}
