package main

import (
	"time"

	"github.com/rpi-ws281x/rpi-ws281x-go"
)

var OFF uint32 = 0x00000000
var GREEN uint32 = 0x00FF0000
var RED uint32 = 0xFF000000
var BLUE uint32 = 0x0000FF00
var WHITE uint32 = 0xFFFFFF00

func main() {
	cOpt := ws2811.ChannelOption{
		GpioPin : 18,
		Invert : false,
		LedCount : 1024,
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

	//snakey(dev, 0, GREEN)

	setLed(dev, 0, 768, GREEN)
	renderLeds(dev)
	time.Sleep(3 * time.Second)
	setLed(dev, 0, 768, OFF)
	renderLeds(dev)
}

func snakey(dev *ws2811.WS2811, channel int, col uint32) {
	for i := range dev.Leds(channel) {
		setLed(dev, channel, i, col)
		renderLeds(dev)
		time.Sleep(time.Millisecond * 100)
		setLed(dev, channel, i, OFF)
	}
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
