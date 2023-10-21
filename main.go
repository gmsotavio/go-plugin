package main

import (
	"go-plugins/atmel"
	"go-plugins/microchip"
)

// Flasher é uma interface que define operações de gravação em microcontroladores.
type Flasher interface {
	Connect()
	WriteData(data string)
	Disconnect()
}

func main() {
	// Escolha a implementação desejada
	microchipImpl := &microchip.MicrochipFlasher{}
	atmelImpl := &atmel.AtmelFlasher{}

	// Use a implementação Microchip
	flash(microchipImpl)

	// Use a implementação Atmel
	flash(atmelImpl)
}

func flash(f Flasher) {
	f.Connect()
	f.WriteData("Data to be flashed")
	f.Disconnect()
}
