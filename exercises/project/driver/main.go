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
	elevTx := make(chan elevator.ElevatorMessage)
	elevRx := make(chan elevator.ElevatorMessage)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollObstructionSwitch(drvObstr)
	go elevio.PollStopButton(drvStop)
	go elevator.FsmPollButtonRequest(drvButtons)
	go bcast.Receiver(15647,elevRx)
	go bcast.Transmitter(15647,elevTx)
	for {
		select {
		case a := <-drvFloors:
			elevator.FsmFloor(a)
		case a := <-drvStop:
			elevator.FsmStop(a)
		case p:= <-elevRx:
			elevator.FsmMessageReceived(p)
		//case s:= <- drvButtons:
			//msg := elevator.ElevatorMessage{"ORDER",int(s.Button),s.Floor}
		//case s:= <- drvFloors:
			//msg := elevator.ElevatorMessage{"FINISHED", s,0}
		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
