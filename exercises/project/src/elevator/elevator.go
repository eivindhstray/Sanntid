package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

type ElevDir int

var Elev Elevator

const (
	Up   ElevDir = 1
	Down         = -1
	Stop         = 0
)

type Elevator struct {
	ElevID       int
	CurrentFloor int
	Dir          ElevDir
	DoorTimer    *time.Timer
	DoorState    bool
	ElevState    variables.ElevatorList
	ElevOnline   int
}

//Update list containing info of elevator. Important to determine cost
func ElevatorListUpdate(ID int, floor int, newDirection ElevDir, connectionStatus int) {
	Elev.ElevState[ID][0] = floor
	Elev.ElevState[ID][1] = int(newDirection)
	Elev.ElevState[ID][2] = connectionStatus
}

//Do we need this?
func ElevatorSetConnectionStatus(connectionStatus int, ID int) {
	Elev.ElevState[ID][2] = connectionStatus
}

//Initialize elevator. Update elevator list.
func ElevatorInit(ID int) {
	if elevio.GetFloor() == -1 {
		elevatorSetDir(Down)
	}
	for elevio.GetFloor() == -1 {
	}
	elevatorSetDir(Stop)
	elevatorSetFloor(elevio.GetFloor())
	Elev.ElevID = ID
	Elev.DoorTimer = time.NewTimer(0)
	ElevatorSetDoorOpenState(false)
	elevio.SetDoorOpenLamp(false)
	ElevatorListUpdate(Elev.ElevID, Elev.CurrentFloor, Elev.Dir, Elev.ElevOnline)
	fmt.Println("Elevator initialized")
}

//Sets new floor for elevator.
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

//Syncronizing ligths to match queue.
func elevatorLightsMatchQueue() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for button := 0; button < variables.N_BUTTON_TYPES; button++ {
			if queueLocal[floor][button] == variables.LOCAL { //|| queueLocal[floor][button] == variables.REMOTE {
				elevio.SetButtonLamp(elevio.ButtonType(button), floor, true)
			} else {
				elevio.SetButtonLamp(elevio.ButtonType(button), floor, false)
			}
		}
	}
}

//Set new direction for elevator.
func elevatorSetDir(newDirection ElevDir) {
	Elev.Dir = newDirection
	elevatorSetMotorDir(newDirection)
}

//Set motor direction.
func elevatorSetMotorDir(newDirection ElevDir) {
	elevio.SetMotorDirection(elevio.MotorDirection(newDirection))
}

//Set new floor for elevator.
func elevatorSetFloor(newFloor int) {
	Elev.CurrentFloor = newFloor
}

//Returns direction of elevator.
func elevatorGetDir() ElevDir {
	return Elev.Dir
}

//Returns elevators floor.
func elevatorGetFloor() int {
	return Elev.CurrentFloor
}

//Set door state for elevator.
func ElevatorSetDoorOpenState(state bool) {
	Elev.DoorState = state
}

//Returns open door state.
func ElevatorGetDoorOpenState() bool {
	return Elev.DoorState
}

//Return elevator struct.
func ElevatorGetElev() Elevator {
	return Elev
}

//---------------------------------------------------------------
//Attempt to make a channel for motor direction.
func ElevatorChannelGetDir(reciever chan<- ElevDir) {
	reciever <- Elev.Dir
}
