package main

import (
	"fmt"

	"./elevator"
	"./elevio"
	"./network/bcast"
	"./variables"
)

func main() {

	elevio.Init("localhost:15657", variables.N_FLOORS)
	network.NetworkInit()

	elevator.ElevatorInit()
	elevator.QueueInit()
	fmt.Println("Initialized")

	// Channels
	drvButtons := make(chan elevio.ButtonEvent)
	drvFloors := make(chan int)
	drvObstr := make(chan bool)
	drvStop := make(chan bool)
<<<<<<< HEAD
	elevTx := make(chan elevator.ElevatorMessage)
=======
	//elevTx := make(chan elevator.ElevatorMessage)
>>>>>>> c0c292c39adbefe5e15283f72078d182ea4dffff
	elevRx := make(chan elevator.ElevatorMessage)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollObstructionSwitch(drvObstr)
	go elevio.PollStopButton(drvStop)
	go elevator.FsmPollButtonRequest(drvButtons)
<<<<<<< HEAD
	go bcast.Receiver(15647,elevRx)
	go bcast.Transmitter(15647,elevTx)
=======
	go bcast.Receiver(15647, elevRx)

>>>>>>> c0c292c39adbefe5e15283f72078d182ea4dffff
	for {
		select {
		case a := <-drvFloors:
			elevator.FsmFloor(a)
		case a := <-drvStop:
			elevator.FsmStop(a)
<<<<<<< HEAD
		case p:= <-elevRx:
			elevator.FsmMessageReceived(p)
		//case s:= <- drvButtons:
			//msg := elevator.ElevatorMessage{"ORDER",int(s.Button),s.Floor}
		//case s:= <- drvFloors:
			//msg := elevator.ElevatorMessage{"FINISHED", s,0}
=======
		case a := <-elevRx:
			elevator.FsmMessageReveived(a)
>>>>>>> c0c292c39adbefe5e15283f72078d182ea4dffff
		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
