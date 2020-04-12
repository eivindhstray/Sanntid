package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloor(newFloor int, dir ElevDir) {

	//decisionAlgorithm(newFloor, elev.dir)
	//Function that updates elevatorList with position and direction
	elevatorSetNewFloor(newFloor)
	if localQueueCheckCurrentFloorSameDir(newFloor, elev.Dir) == true {
		fsmDoorState()
	}
	localQueueRemoveOrder(newFloor, dir)
	elevatorLightsMatchQueue()
	elevatorSetDir(localQueueReturnElevDir(newFloor, elev.Dir))

}

func fsmOnButtonRequest(buttonPush elevio.ButtonEvent, cabCall bool) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", buttonPush)

	//----------Part that needs work ---------------
	if !cabCall{
		remoteQueueRecieveOrder(buttonPush)
		decisionAlgorithm()
	}else{
		localQueueRecieveOrder(buttonPush)
	}
	remoteQueuePrint()
	localQueuePrint()
	//----------------------------------------------

	elevatorLightsMatchQueue()
	elev = ElevatorGetElev()
	previousDirection := elev.Dir
	if elev.Dir == Stop && !ElevatorGetDoorOpenState() {
		if buttonPush.Floor == elev.currentFloor && elev.Dir == Stop {
			FsmFloor(elev.currentFloor,previousDirection)
		}
		if ElevatorGetDoorOpenState() == false {
			elevatorSetDir(localQueueReturnElevDir(elev.currentFloor, elev.Dir))
		}
	}
}

func FsmMessageReceivedHandler(msg variables.ElevatorMessage, ID int) {
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
			if msgID == ID {
				fsmOnButtonRequest(event, true)
			} else {
				fmt.Println("cabcall other elev")
			}
		} else {
			fsmOnButtonRequest(event, false)
		}
	case "FLOOR":
		if msgID == ID {
			fmt.Print("Floor\v%q", msgID)
			FsmFloor(floor,ElevDir(dir))
		}
		ElevatorListUpdate(msgID, floor)
	case "ALIVE":
		fmt.Println("Alive from", msgID)
	default:
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}

func fsmDoorState() {
	fmt.Print("door")
	elevatorSetMotorDir(Stop)
	ElevatorSetDoorOpenState(true)
	elevio.SetDoorOpenLamp(true)
	elev.doorTimer.Stop()
	elev.doorTimer.Reset(variables.DOOROPENTIME * time.Second)
	<-elev.doorTimer.C
	elevio.SetDoorOpenLamp(false)
	ElevatorSetDoorOpenState(false)

}

//From project destription in the course embedded systems
func FsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	elev = ElevatorGetElev()
	ElevatorInit(elev.ElevID)
	LocalQueueInit()
	elevatorLightsMatchQueue()
}

