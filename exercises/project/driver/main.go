package main

import (
	"fmt"

	"./elevator"
	"./elevio"
	"./variables"
	"./network/bcast"
)

func main() {

	elevio.Init("localhost:15657", variables.N_FLOORS)

	elevator.ElevatorInit()
	elevator.QueueInit()
	fmt.Println("Initialized")

	// Channels
	drvButtons := make(chan elevio.ButtonEvent)
	drvFloors := make(chan int)
	drvObstr := make(chan bool)
	drvStop := make(chan bool)
	elevTx := make(chan variables.ElevatorMessage)
	elevRx := make(chan variables.ElevatorMessage)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollObstructionSwitch(drvObstr)
	go elevio.PollStopButton(drvStop)
	

	go elevator.FsmPollButtonRequest(drvButtons)

	for {
		select {
		case a := <-drvFloors:
			elevator.FsmFloor(a)
		case a := <-drvStop:
			elevator.FsmStop(a)

		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
