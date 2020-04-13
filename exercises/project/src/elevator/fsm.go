package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloor(newFloor int, dir ElevDir, msgID int) {

	//decisionAlgorithm(newFloor, elev.dir)
	//Function that updates elevatorList with position and direction
	ElevatorListUpdate(msgID, newFloor)
	fmt.Println(Elev.ElevState,"   ",msgID,"   ", newFloor)
	if msgID == Elev.ElevID{
		elevatorSetNewFloor(newFloor)
	}
	if localQueueCheckCurrentFloorSameDir(newFloor, Elev.Dir) == true {
		fsmStartDoorState(Elev.DoorTimer)
	}
	localQueueRemoveOrder(newFloor, dir)
	elevatorLightsMatchQueue()
	if msgID == Elev.ElevID{
		elevatorSetDir(localQueueReturnElevDir(newFloor, Elev.Dir))
	}
	remoteQueuePrint()
	localQueuePrint()

}

func fsmOnButtonRequest(buttonPush elevio.ButtonEvent, cabCall bool) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", buttonPush)

	//----------Part that needs work ---------------
	if !cabCall {
		remoteQueueRecieveOrder(buttonPush)
		decisionAlgorithm(buttonPush)
	} else {
		localQueueRecieveOrder(buttonPush)
	}
	//----------------------------------------------

	elevatorLightsMatchQueue()
	elev := ElevatorGetElev()
	previousDirection := elev.Dir
	if elev.Dir == Stop && !ElevatorGetDoorOpenState() {
		if buttonPush.Floor == elev.currentFloor && elev.Dir == Stop {
			FsmFloor(elev.currentFloor, previousDirection, elev.ElevID)
		}
		if ElevatorGetDoorOpenState() == false {
			elevatorSetDir(localQueueReturnElevDir(elev.currentFloor, elev.Dir))
		}
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
	switch msgType {
	case "ORDER":
		if button == 2 {
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
		FsmFloor(floor, ElevDir(dir),msgID)
	case "ALIVE":
		fmt.Println("Alive from", msgID)
	default:
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}

func fsmStartDoorState(doorTimer *time.Timer) {
	fmt.Print("door")
	elevatorSetMotorDir(Stop)
	ElevatorSetDoorOpenState(true)
	elevio.SetDoorOpenLamp(true)
	doorTimer.Reset(variables.DOOROPENTIME * time.Second)
}

func FsmExitDoorState(doorTimer *time.Timer){
	doorTimer.Stop()
	ElevatorSetDoorOpenState(false)
	elevio.SetDoorOpenLamp(false)
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
