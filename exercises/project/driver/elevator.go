package main

import (
	"fmt"

	"./elevio"
)

type ElevDir int

const (
	Up   ElevDir = 1
	Down         = -1
	Stop         = 0
)

type Elevator struct {
	currentFloor int
	dir          ElevDir
}

var elevator Elevator

func elevatorInit() {
	if elevio.GetFloor() == -1 {
		elevatorSetDir(Down)
	}
	for elevio.GetFloor() == -1 {
	}
	elevatorSetDir(Stop)
	elevatorSetFloor(elevio.GetFloor())

	fmt.Println("Elevator initialized")
}

func elevatorSetNewFloor(newFloor int) {

	elevatorSetFloor(newFloor)
	switch newFloor {
	case N_FLOORS - 1:
		elevatorSetDir(Down)
		break
	case 0:
		elevatorSetDir(Up)
		break
	}
}

func elevatorLightsMatchQueue() {
	for floor := 0; floor < N_FLOORS; floor++ {
		for button := 0; button < N_BUTTON_TYPES; button++ {
			if queue[floor][button] == true {
				elevio.SetButtonLamp(elevio.ButtonType(button), floor, true)
			} else {
				elevio.SetButtonLamp(elevio.ButtonType(button), floor, false)
			}
		}
	}
}

func elevatorSetDir(newDirection ElevDir) {
	elevator.dir = newDirection
	elevatorSetMotorDir(newDirection)
}

func elevatorSetMotorDir(newDirection ElevDir) {
	elevio.SetMotorDirection(elevio.MotorDirection(newDirection))
}

func elevatorSetFloor(newFloor int) {
	elevator.currentFloor = newFloor
}

func elevatorGetDir() ElevDir {
	return elevator.dir
}

func elevatorGetFloor() int {
	return elevator.currentFloor
}

func elevatorPrint() {

}
