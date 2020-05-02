package main

import (
	"fmt"
	"time"
	"log"

	"github.com/rakyll/portmidi"
)

func main() {
	portmidi.Initialize()
	
	numDevices := portmidi.CountDevices()
	inputId := portmidi.DefaultInputDeviceID()
	outputId := portmidi.DefaultOutputDeviceID()
	
	in, inErr := portmidi.NewInputStream(inputId, 1024)
	if inErr != nil {
		log.Fatal(inErr)
	}
	out, outErr := portmidi.NewOutputStream(outputId, 1024, 0)
	if outErr != nil {
		log.Fatal(outErr)
	}

	fmt.Println("NUMBER OF DEVICES: ", numDevices)
	fmt.Println("INPUT ID: ", inputId)
	fmt.Println("OUTPUT ID: ", outputId)
	fmt.Println(portmidi.Info(inputId))
	fmt.Println(portmidi.Info(outputId))

	out.WriteShort(0x90, 60, 100)
	out.WriteShort(0x90, 64, 100)
	out.WriteShort(0x90, 67, 100)
	
	time.Sleep(2 * time.Second)

	out.WriteShort(0x80, 60, 100)
	out.WriteShort(0x80, 64, 100)
	out.WriteShort(0x80, 67, 100)

	// WriteCMaj(out)

	in.Close()
	out.Close()
	portmidi.Terminate()
	fmt.Println("DONE")
}

// func WriteCMaj(out *portmidi.Stream) {
// 	out.WriteShort(0x90, 60, 100)
// 	out.WriteShort(0x90, 64, 100)
// 	out.WriteShort(0x90, 67, 100)
	
// 	time.Sleep(2 * time.Second)

// 	out.WriteShort(0x80, 60, 100)
// 	out.WriteShort(0x80, 64, 100)
// 	out.WriteShort(0x80, 67, 100)
// }
