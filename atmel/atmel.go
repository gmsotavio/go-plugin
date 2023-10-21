package atmel

import "fmt"

type AtmelFlasher struct{}

func (a *AtmelFlasher) Connect() {
	fmt.Println("Atmel: Connecting...")
}

func (a *AtmelFlasher) WriteData(data string) {
	fmt.Printf("Atmel: Writing data '%s'\n", data)
}

func (a *AtmelFlasher) Disconnect() {
	fmt.Println("Atmel: Disconnecting...")
}

