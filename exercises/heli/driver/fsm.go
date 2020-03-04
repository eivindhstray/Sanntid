package main

import (
	"fmt"
	"time"

	"./elevio"
)

func fsmFloor(newFloor int) {
	for i := 0; i < 2; i++ {
		if queueCheckCurrentFloorSameDir(newFloor, elevatorGetDir()) {
			elevatorSetMotorDir(Stop)
			fsmDoorState()
			queueRemoveOrder(newFloor, elevatorGetDir())
			elevatorLightsMatchQueue()
		}

		elevatorSetDir(queueReturnElevDir(newFloor, elevatorGetDir()))
		// Print eleator stuff here
	}
}

func fsmPollButtonRequest(drvButtons chan elevio.ButtonEvent) {
	for {
		fsmOnButtonRequest(<-drvButtons)
	}
}

func fsmOnButtonRequest(a elevio.ButtonEvent) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", a)

	if a.Floor == elevatorGetFloor() && elevatorGetDir() == Stop {
		fsmDoorState()
		return
	}

	queueRecieveOrder(a)
	elevatorLightsMatchQueue()

	if elevatorGetDir() == Stop {
		elevatorSetDir(queueReturnElevDir(elevatorGetFloor(), elevatorGetDir()))
	}
}

func fsmDoorState() {
	fmt.Print("Door state")
	elevio.SetDoorOpenLamp(true)
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	elevio.SetDoorOpenLamp(false)
}

func fsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	elevatorInit()
	queueInit()
	elevatorLightsMatchQueue()
}
