package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"./elevator"
	"./elevio"
	"./network/bcast"
	"./network/localip"
	"./variables"
)

func main() {

	cmd := os.Args[1]
	ElevatorID, err := strconv.Atoi(os.Args[2])
	fmt.Println(ElevatorID)
	time.Sleep(1 * time.Second)

	if err != nil {
		panic(err)
	}
	elevio.Init("localhost:"+cmd, variables.N_FLOORS)
	//go run main.go portnr id
	elevator.ElevatorInit(ElevatorID)
	fmt.Println("Initialized")

	elevator.LocalQueueInit()

	var id string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.Parse()

	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}

	// Channels
	drvButtons := make(chan elevio.ButtonEvent)
	drvFloors := make(chan int)
	drvStop := make(chan bool)
	elevTx := make(chan variables.ElevatorMessage)
	elevRx := make(chan variables.ElevatorMessage)
	timeOut := time.NewTimer(0)
	DoorTimer := elevator.Elev.DoorTimer
	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollStopButton(drvStop)
	go bcast.Receiver(15648, elevRx)
	go bcast.Transmitter(15648, elevTx)

	for {
		select {
		case atFloor := <-drvFloors:
			elevator.ElevatorListUpdate(ElevatorID, atFloor, elevator.Elev.Dir, elevator.Elev.ElevOnline)
			elevator.FsmFloor(atFloor, elevator.Elev.Dir)
			msg := variables.ElevatorMessage{ElevatorID, "FLOOR", -1, atFloor, int(elevator.Elev.Dir), elevator.Elev.ElevState}
			//fmt.Printf("elevstates%q\n", elevator.Elev.ElevState)
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
		case stop := <-drvStop:
			elevator.FsmStop(stop)
		case elevatorMessageReceived := <-elevRx:
			//fmt.Print("Message nr", elevatorMessageReceived.ElevID)
			elevator.FsmMessageReceivedHandler(elevatorMessageReceived, ElevatorID)
			if !elevator.CheckQueueEmpty(variables.LOCAL) {
				timeOut.Reset(variables.FAULT_TIME * time.Second)
			} else {
				timeOut.Stop()
			}
		case buttonCall := <-drvButtons:
			if buttonCall.Button == elevator.Cab {
				elevator.FsmOnButtonRequest(buttonCall, true)
			} else {
				elev := elevator.ElevatorGetElev()
				msg := variables.ElevatorMessage{ElevatorID, "ORDER", int(buttonCall.Button), buttonCall.Floor, int(elev.Dir), elev.ElevState}
				elevTx <- msg
				elevTx <- msg
				elevTx <- msg
				elevTx <- msg
			}

		case <-timeOut.C:
			fmt.Printf("Timer fired")
			elevator.ElevatorSetConnectionStatus(variables.NEW_FLOOR_TIMEOUT_PENALTY, ElevatorID)
			elev := elevator.ElevatorGetElev()
			msg := variables.ElevatorMessage{ElevatorID, "FAULTY_MOTOR", -1, -1, int(elev.Dir), elev.ElevState}
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg

		case <-DoorTimer.C:
			elevator.FsmExitDoorState(elevator.Elev.DoorTimer)
		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
