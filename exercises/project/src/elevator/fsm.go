package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

func FsmFloor(newFloor int) {

	//decisionAlgorithm(newFloor, elev.dir)
	elevatorSetNewFloor(newFloor)
	if localQueueCheckCurrentFloorSameDir(newFloor, elev.dir) {
		elevatorSetMotorDir(Stop)
		localQueueRemoveOrder(newFloor, elev.dir)
		elevatorLightsMatchQueue()
		fsmDoorState()
	}
	elevatorSetDir(localQueueReturnElevDir(newFloor, elev.dir))

}

func fsmOnButtonRequest(a elevio.ButtonEvent) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", a)
	//remoteQueueRecieveOrder(a)
	//decisionAlgorithm(elev.currentFloor, elev.dir)
	localQueueRecieveOrder(a)
	elevatorLightsMatchQueue()
	elev = ElevatorGetElev()
	if elev.dir == Stop {
		if a.Floor == elev.currentFloor && elev.dir == Stop {
			//fsmDoorState()
			FsmFloor(elev.currentFloor)
		}
		if ElevatorGetDoorOpenState() == false {
			elevatorSetDir(localQueueReturnElevDir(elev.currentFloor, elev.dir))
		}
	}
}

func FsmMessageReceivedHandler(msg ElevatorMessage, ID string) {
	//sync the new message with queue
	fmt.Println("received a message")
	msgType := msg.MessageType
	msgID := msg.ElevID
	floor := msg.Floor
	button := msg.Button
	event := elevio.ButtonEvent{floor, elevio.ButtonType(button)}
	switch msgType {
	case "ORDER":
		fmt.Println(msgID+"+"+ID, "+", button)
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
		FsmFloor(floor)
	case "ALIVE":
		fmt.Println("Alive from", msgID)
	default:
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}

func fsmDoorState() {
	fmt.Print("Door state")
	ElevatorSetDoorOpenState(true)
	elevio.SetDoorOpenLamp(true)
	elev.doorTimer.Reset(variables.DOOROPENTIME * time.Second)
	<-elev.doorTimer.C
	fmt.Print("DoorState over")
	elevio.SetDoorOpenLamp(false)
	ElevatorSetDoorOpenState(false)
}

//From project destription in the course embedded systems
func FsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	ElevatorInit()
	LocalQueueInit()
	elevatorLightsMatchQueue()
}
