package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)


func FsmFloor(newFloor int) {

	elevatorSetNewFloor(newFloor)
	if localQueueCheckCurrentFloorSameDir(newFloor, elev.dir) {
		elevatorSetMotorDir(Stop)
		fsmDoorState()
		localQueueRemoveOrder(newFloor, elev.dir)
		elevatorLightsMatchQueue()

	}
	elevatorSetDir(localQueueReturnElevDir(newFloor, elev.dir))

}

func fsmOnButtonRequest(a elevio.ButtonEvent) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", a)
	localQueueRecieveOrder(a)
	elevatorLightsMatchQueue()
	elev = ElevatorGetElev()

	if elev.dir == Stop {
		if a.Floor == elev.currentFloor && elev.dir == Stop {
			fsmDoorState()
			FsmFloor(elev.currentFloor)
		}
		if ElevatorGetDoorState() == false{ 
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
	if msgType == "ORDER"{
		fmt.Println(msgID +"+"+ ID,"+",button)
		if button == 2{
			if msgID == ID {
				fsmOnButtonRequest(event)
			}else{
				fmt.Println("cabcall other elev")
			}
		}else{
			fsmOnButtonRequest(event)
		}
	} else if msgType == "FLOOR" {
		FsmFloor(floor)
	} else {
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}


func fsmDoorState() {
	fmt.Print("Door state")
	ElevatorSetDoorState(true)
	elevio.SetDoorOpenLamp(true)
	elev.doorTimer.Reset(variables.DOOROPENTIME * time.Second)
	<-elev.doorTimer.C
	ElevatorSetDoorState(false)
	elevio.SetDoorOpenLamp(false)

}

//From project destription in the course embedded systems
func FsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	ElevatorInit()
	LocalQueueInit()
	elevatorLightsMatchQueue()
}

