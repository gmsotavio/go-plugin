package microchip

import "fmt"

type MicrochipFlasher struct{}

func (m *MicrochipFlasher) Connect() {
	fmt.Println("Microchip: Connecting...")
}

func (m *MicrochipFlasher) WriteData(data string) {
	fmt.Printf("Microchip: Writing data '%s'\n", data)
}

func (m *MicrochipFlasher) Disconnect() {
	fmt.Println("Microchip: Disconnecting...")
}

