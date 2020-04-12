package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloor(newFloor int) {

	//decisionAlgorithm(newFloor, elev.dir)
	//Function that updates elevatorList with position and direction
	elevatorSetNewFloor(newFloor)
	if localQueueCheckCurrentFloorSameDir(newFloor, elev.dir) == true {
		fsmDoorState()
	}
	elevatorSetDir(localQueueReturnElevDir(newFloor, elev.dir))

}

func fsmOnButtonRequest(buttonPush elevio.ButtonEvent) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", buttonPush)

	//----------Part that needs work ---------------
	remoteQueueRecieveOrder(buttonPush)
	decisionAlgorithm()
	remoteQueuePrint()
	localQueuePrint()
	//----------------------------------------------

	elevatorLightsMatchQueue()
	elev = ElevatorGetElev()
	if elev.dir == Stop && !ElevatorGetDoorOpenState() {
		if buttonPush.Floor == elev.currentFloor && elev.dir == Stop {
			FsmFloor(elev.currentFloor)
		}
		if ElevatorGetDoorOpenState() == false {
			elevatorSetDir(localQueueReturnElevDir(elev.currentFloor, elev.dir))
		}
	}
}

func FsmMessageReceivedHandler(msg variables.ElevatorMessage, ID int) {
	//sync the new message with queue
	fmt.Println("received a message")
	msgType := msg.MessageType
	msgID := msg.ElevID
	floor := msg.Floor
	button := msg.Button
	event := elevio.ButtonEvent{floor, elevio.ButtonType(button)}
	switch msgType {
	case "ORDER":
		if button == 2 {
			if msgID == ID {
				fsmOnButtonRequest(event)
			} else {
				fmt.Println("cabcall other elev")
			}
		} else {
			fsmOnButtonRequest(event)
		}
	case "FLOOR":
		if msgID == ID {
			fmt.Print("Floor\v%q", msgID)
			FsmFloor(floor)
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
	fsmRemoveOrderHandler(elev)

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

func fsmRemoveOrderHandler(elev Elevator) {
	localQueueRemoveOrder(elev.currentFloor, elev.dir)
	elevatorLightsMatchQueue()
}
