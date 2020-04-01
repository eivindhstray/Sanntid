package elevator

import (
	"fmt"
	"time"

	"../elevio"
	"../variables"
)

//seems like there is a bug related to cab calls. The elevator sometimes go
//out of bounds

func FsmFloor(newFloor int) {
	for i := 0; i < 2; i++ {
		if queueCheckCurrentFloorSameDir(newFloor, elevatorGetDir()) {
			elevatorSetMotorDir(Stop)
			fsmDoorState()
			queueRemoveOrder(newFloor, elevatorGetDir())
			backupSync()
			elevatorLightsMatchQueue()
		}

		elevatorSetDir(queueReturnElevDir(newFloor, elevatorGetDir()))
		// Print eleator stuff here
	}
}

func FsmPollButtonRequest(drvButtons chan elevio.ButtonEvent) {
	for {
		fsmOnButtonRequest(<-drvButtons)
	}
}

func fsmOnButtonRequest(a elevio.ButtonEvent) {
	fmt.Print("New order recieved")
	fmt.Printf("%+v\n", a)

	if a.Floor == elevatorGetFloor() && elevatorGetDir() == Stop {
		fsmDoorState()
		return
	}

	queueRecieveOrder(a)
	elevatorLightsMatchQueue()
	backupSync()

	if elevatorGetDir() == Stop {
		elevatorSetDir(queueReturnElevDir(elevatorGetFloor(), elevatorGetDir()))
	}
}

func FsmMessageReceived(msg ElevatorMessage) {
	//sync the new message with queue
	msgType := msg.MessageType
	button := int(msg.Button)
	floor := msg.Floor
	if msgType == "ORDER" {
		queueSet(floor, int(button))
	} else if msgType == "FINISHED" {
		queuePop(floor, int(button))
	} else {
		fmt.Print("invalid message")
	}
	elevatorLightsMatchQueue()

}

func FsmMessageTransmit(msgType string, floor int, button int) {

}

func fsmDoorState() {
	fmt.Print("Door state")
	elevio.SetDoorOpenLamp(true)
	timer1 := time.NewTimer(variables.DOOROPENTIME * time.Second)
	<-timer1.C
	elevio.SetDoorOpenLamp(false)
}

//From project destription in the course embedded systems
func FsmStop(a bool) {
	fmt.Print("Stop state")
	fmt.Printf("%+v\n", a)
	ElevatorInit()
	QueueInit()
	elevatorLightsMatchQueue()
}
