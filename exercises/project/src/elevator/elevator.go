package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

type ElevDir int

var elev Elevator

const (
	Up   ElevDir = 1
	Down         = -1
	Stop         = 0
)

type Elevator struct {
	ElevID 		 int
	currentFloor int
	dir          ElevDir
	doorTimer    *time.Timer
	DoorState    bool
	ElevState	 variables.ElevatorList
}

func ElevatorListUpdate(ID int,floor int) {
	elev.ElevState[ID][0] = floor
}

func ElevatorInit(ID int) {
	if elevio.GetFloor() == -1 {
		elevatorSetDir(Down)
	}
	for elevio.GetFloor() == -1 {
	}
	elevatorSetDir(Stop)
	elevatorSetFloor(elevio.GetFloor())
	elev.doorTimer = time.NewTimer(0)
	elev.ElevID = ID

	fmt.Println("Elevator initialized")
}

func elevatorSetNewFloor(newFloor int) {

	elevatorSetFloor(newFloor)
	elevio.SetFloorIndicator(newFloor)
	switch newFloor {
	case variables.N_FLOORS - 1:
		elevatorSetDir(Down)
		break
	case 0:
		elevatorSetDir(Up)
		break
	}
}

func elevatorLightsMatchQueue() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queueLocal[floor][button] == variables.LOCAL || queueLocal[floor][button] == variables.REMOTE {
				elevio.SetButtonLamp(elevio.ButtonType(button), floor, true)
			} else {
				elevio.SetButtonLamp(elevio.ButtonType(button), floor, false)
			}
		}
	}
}

func elevatorSetDir(newDirection ElevDir) {
	elev.dir = newDirection
	elevatorSetMotorDir(newDirection)
}

func elevatorSetMotorDir(newDirection ElevDir) {
	elevio.SetMotorDirection(elevio.MotorDirection(newDirection))
}

func elevatorSetFloor(newFloor int) {
	elev.currentFloor = newFloor
}

func elevatorGetDir() ElevDir {
	return elev.dir
}

func elevatorGetFloor() int {
	return elev.currentFloor
}

func ElevatorSetDoorOpenState(state bool) {
	elev.DoorState = state
}

func ElevatorGetDoorOpenState() bool {
	return elev.DoorState
}

func ElevatorGetElev() Elevator {
	return elev
}
