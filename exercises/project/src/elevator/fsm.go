package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloorMessage(newFloor int, dir ElevDir, msgID int) {
	BackupSyncQueue()
	if msgID != Elev.ElevID{
		QueueRemoveOrder(newFloor, dir)
	}
	elevatorLightsMatchQueue()
	QueuePrintRemote()
	QueuePrintLocal()
	ElevatorListUpdate(msgID, newFloor, Elev.Dir, Elev.ElevOnline)

}

func FsmFloor(newFloor int, dir ElevDir){
	elevatorSetNewFloor(newFloor)
	elevatorLightsMatchQueue()
	if QueueCheckCurrentFloorSameDir(newFloor, Elev.Dir) {
		fmt.Println("stopping")
		QueueRemoveOrder(newFloor,dir)
		fsmStartDoorState(Elev.DoorTimer)
	}
	if !Elev.DoorState {
		elevatorSetDir(QueueReturnElevDir(newFloor, Elev.Dir))
	}
	if CheckQueueEmpty(variables.LOCAL){
		elevatorSetDir(Stop)
	}
	
}

func fsmOnButtonRequest(buttonPush elevio.ButtonEvent, cabCall bool) {

	if !cabCall {
		QueueRecieveOrderRemote(buttonPush)
		decisionAlgorithm()
	} else {
		QueueRecieveOrderLocal(buttonPush)
	}
	elevatorLightsMatchQueue()
	if buttonPush.Floor == Elev.CurrentFloor && QueueCheckCurrentFloorSameDir(Elev.CurrentFloor,Elev.Dir){
		FsmFloor(Elev.CurrentFloor, Elev.Dir)
	}
	if !ElevatorGetDoorOpenState() {
		elevatorSetDir(QueueReturnElevDir(Elev.CurrentFloor, Elev.Dir))
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
		FsmFloorMessage(floor, ElevDir(dir), msgID)
	case "FAULTY_MOTOR":
		ElevatorSetConnectionStatus(variables.NEW_FLOOR_TIMEOUT_PENALTY, msgID)
		if msgID != LocalID && Elev.Dir == Stop {
			QueueMakeRemoteLocal()
			FsmFloor(Elev.CurrentFloor, Elev.Dir)
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
	FsmFloor(Elev.CurrentFloor, Elev.Dir)
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
