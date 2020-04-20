package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloorMessage(newFloor int, dir ElevDir, msgID int) {
	QueueRemoveOrder(newFloor, dir, msgID)
	elevatorLightsMatchQueue()
	ElevatorFloorUpdate(msgID, newFloor)
	ElevatorSetConnectionStatus(msgID, Elev.ElevOnline)
}

func FsmFloor(newFloor int, dir ElevDir) {
	elevatorSetNewFloor(newFloor)
	elevatorLightsMatchQueue()
	if queueCheckCurrentFloorSameDir(newFloor, Elev.Dir) {
		fsmStartDoorState(Elev.DoorTimer)
	}
	if !Elev.DoorOpen {
		elevatorSetDir(queueReturnElevDir(newFloor, Elev.Dir))
	}
}

func FsmOnButtonRequest(buttonPush elevio.ButtonEvent, cabCall bool) {
	if buttonPush.Floor == Elev.CurrentFloor && queueCheckCurrentFloorSameDir(Elev.CurrentFloor, Elev.Dir) && Elev.Dir == Stop {
		FsmFloor(Elev.CurrentFloor, Elev.Dir)
	} else if !cabCall {
		QueueRecieveOrderRemote(buttonPush)
		bestElev := DecisionChooseElevator(buttonPush)
		if bestElev == Elev.ElevID {
			QueueRecieveOrderLocal(buttonPush)
		}
	} else {
		QueueRecieveOrderLocal(buttonPush)
	}
	elevatorLightsMatchQueue()
	if !Elev.DoorOpen {
		elevatorSetDir(queueReturnElevDir(Elev.CurrentFloor, Elev.Dir))
	}

}

func FsmMessageReceivedHandler(msg variables.ElevatorMessage, LocalID int) {
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
			if button == int(Cab) {
				FsmOnButtonRequest(event, true)
			}
		} else {
			FsmOnButtonRequest(event, false)
		}
	case "FLOOR":
		FsmFloorMessage(floor, ElevDir(dir), msgID)
	case "NOT_RESPONDING":
		ElevatorSetConnectionStatus(msgID, variables.ELEV_OFFLINE)
		queueMakeRemoteLocal()
		if Elev.Dir == Stop {
			FsmFloor(Elev.CurrentFloor, Elev.Dir)
		}
	default:
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}

func fsmStartDoorState(doorTimer *time.Timer) {
	fmt.Print("door")
	Elev.DoorOpen = true
	elevio.SetDoorOpenLamp(true)
	elevio.SetMotorDirection(Stop)
	QueueRemoveOrder(Elev.CurrentFloor, Elev.Dir, Elev.ElevID)
	doorTimer.Reset(variables.DOOROPENTIME * time.Second)
}

func FsmExitDoorState(doorTimer *time.Timer) {
	doorTimer.Stop()
	Elev.DoorOpen = false
	elevio.SetDoorOpenLamp(false)
	FsmFloor(Elev.CurrentFloor, Elev.Dir)
}

//From project destription in the course embedded systems
func FsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	ElevatorInit(Elev.ElevID)
	LocalQueueInit()
	elevatorLightsMatchQueue()
}
