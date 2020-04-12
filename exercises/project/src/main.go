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
	"./network/peers"
	"./variables"
	//"./watchdog"
)

func main() {

	cmd := os.Args[1]
	ElevatorID, err := strconv.Atoi(os.Args[2])

	//variables.ElevatorID = ElevatorID
	//^^ Need ElevID to be an int for cost to function properly
	fmt.Println(ElevatorID)
	time.Sleep(1 * time.Second)

	if err != nil {
		panic(err)
	}
	elevio.Init("localhost:"+cmd, variables.N_FLOORS)
	//go run main.go portnr id

	elevator.ElevatorInit(ElevatorID)
	elevator.LocalQueueInit()
	//watchdog.WatchDogInit()
	fmt.Println("Initialized")

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
	peerUpdateCh := make(chan peers.PeerUpdate)
	peerTxEnable := make(chan bool)
	//watchDogTimeOut := make(chan bool)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollStopButton(drvStop)
	go bcast.Receiver(15648, elevRx)
	go bcast.Transmitter(15648, elevTx)
	go peers.Transmitter(15647, id, peerTxEnable)
	go peers.Receiver(15647, peerUpdateCh)
	//go watchdog.WatchDogTimeNSeconds(watchDogTimeOut)

	for {
		select {
		case atFloor := <-drvFloors:
			elevator.ElevatorListUpdate(ElevatorID, atFloor)
			elev := elevator.ElevatorGetElev()
			msg := variables.ElevatorMessage{ElevatorID, "FLOOR", -1, atFloor,int(elev.Dir), elev.ElevState}
			fmt.Printf("elevstates%q\n", elev.ElevState)
			elevTx <- msg
		case stop := <-drvStop:
			elevator.FsmStop(stop)
		case messageReceived := <-elevRx:
			elevator.FsmMessageReceivedHandler(messageReceived, ElevatorID)
		case buttonCall := <-drvButtons:
			elev := elevator.ElevatorGetElev()
			msg := variables.ElevatorMessage{ElevatorID, "ORDER", int(buttonCall.Button), buttonCall.Floor,int(elev.Dir), elev.ElevState}
			elevTx <- msg
		case newPeerEvent := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", newPeerEvent.Peers)
			fmt.Printf("  New:      %q\n", newPeerEvent.New)
			fmt.Printf("  Lost:     %q\n", newPeerEvent.Lost)
		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
