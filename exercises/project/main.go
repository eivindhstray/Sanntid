package main

import (
	"fmt"


	
	"./driver/elevio"
	"./driver/fsm"
	"./driver/queue"
	"./driver/elevator"
	"./network/bcast"
)

func main() {

	elevio.Init("localhost:15657", N_FLOORS)

	elevatorInit()
	queueInit()
	fmt.Println("Initialized")

	// Channels
	drvButtons := make(chan elevio.ButtonEvent)
	drvFloors := make(chan int)
	drvObstr := make(chan bool)
	drvStop := make(chan bool)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollObstructionSwitch(drvObstr)
	go elevio.PollStopButton(drvStop)

	go fsmPollButtonRequest(drvButtons)

	for {
		select {
		case a := <-drvFloors:
			fsmFloor(a)
		case a := <-drvStop:
			fsmStop(a)

		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
