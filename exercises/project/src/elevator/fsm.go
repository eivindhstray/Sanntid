package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloor(newFloor int, dir ElevDir, msgID int, cabCall bool) {
	if msgID == Elev.ElevID {
		elevatorSetNewFloor(newFloor)
		elevatorLightsMatchQueue()
	}
	if localQueueCheckCurrentFloorSameDir(newFloor, Elev.Dir) == true {
		fsmStartDoorState(Elev.DoorTimer)
	}
	if !cabCall {
		localQueueRemoveOrder(newFloor, dir)
		elevatorLightsMatchQueue()
	}
	if !Elev.DoorState {
		elevatorSetDir(localQueueReturnElevDir(newFloor, Elev.Dir))
	}
	remoteQueuePrint()
	localQueuePrint()
	ElevatorListUpdate(msgID, newFloor, Elev.Dir, Elev.ElevOnline)

}

func fsmOnButtonRequest(buttonPush elevio.ButtonEvent, cabCall bool) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", buttonPush)

	if !cabCall {
		remoteQueueRecieveOrder(buttonPush)
		decisionAlgorithm(buttonPush)
	} else {
		localQueueRecieveOrder(buttonPush)
	}

	elevatorLightsMatchQueue()
	fmt.Println("Direction: ", Elev.Dir)

	if buttonPush.Floor == Elev.CurrentFloor && elevatorGetDir() == Stop {
		FsmFloor(Elev.CurrentFloor, Elev.Dir, Elev.ElevID, cabCall)
	}
	if !ElevatorGetDoorOpenState() {
		elevatorSetDir(localQueueReturnElevDir(Elev.CurrentFloor, Elev.Dir))
	}

}

func FsmMessageReceivedHandler(msg variables.ElevatorMessage, LocalID int) {
	//sync the new message with queue
	fmt.Println("received a message")
	msgType := msg.MessageType
	msgID := msg.ElevID
	floor := msg.Floor
	dir := msg.Dir
	button := msg.Button
	event := elevio.ButtonEvent{floor, elevio.ButtonType(button)}
	cabCall := false
	if button == 2 {
		cabCall = true
	}
	switch msgType {
	case "ORDER":
		if cabCall {
			if msgID == LocalID {
				fsmOnButtonRequest(event, true)
			} else {
				fmt.Println("cabcall other elev")
			}
		} else {
			fsmOnButtonRequest(event, false)
		}
	case "FLOOR":
		if msgID == LocalID {
			fmt.Print("Floor\v%q", msgID)
		}
		for elevio.GetFloor() == -1 {
		}
		elevatorSetDir(Stop)
		FsmFloor(floor, ElevDir(dir), msgID, cabCall)
	case "FAULTY_MOTOR":
		if msgID == LocalID {
			ElevatorSetConnectionStatus(variables.NEW_FLOOR_TIMEOUT_PENALTY, msgID)
		}
		if msgID != LocalID && Elev.Dir == Stop {
			FsmFloor(Elev.CurrentFloor, Elev.Dir, LocalID, false)
		}
	default:
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}

func fsmStartDoorState(doorTimer *time.Timer) {
	fmt.Print("door")
	elevatorSetDir(Stop)
	ElevatorSetDoorOpenState(true)
	elevio.SetDoorOpenLamp(true)
	doorTimer.Reset(variables.DOOROPENTIME * time.Second)
}

func FsmExitDoorState(doorTimer *time.Timer) {
	doorTimer.Stop()
	ElevatorSetDoorOpenState(false)
	elevio.SetDoorOpenLamp(false)
	FsmFloor(Elev.CurrentFloor, Elev.Dir, Elev.ElevID, false)
}

//From project destription in the course embedded systems
func FsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	elev := ElevatorGetElev()
	ElevatorInit(elev.ElevID)
	LocalQueueInit()
	elevatorLightsMatchQueue()
}
