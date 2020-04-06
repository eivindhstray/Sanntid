package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"./elevator"
	"./elevio"
	"./network/bcast"
	"./network/localip"
	"./network/peers"
	"./variables"

)

func main() {

	cmd := os.Args[1]
	ElevatorID := os.Args[2] 
	fmt.Println(cmd)
	elevio.Init("localhost:"+cmd, variables.N_FLOORS)
	//go run main.go portnr

	elevator.ElevatorInit()
	elevator.QueueInit()
	//elevator.backupInit()
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

	peerUpdateCh := make(chan peers.PeerUpdate)
	peerTxEnable := make(chan bool)
	go peers.Transmitter(15647, id, peerTxEnable)
	go peers.Receiver(15647, peerUpdateCh)

	// Channels
	drvButtons := make(chan elevio.ButtonEvent)
	drvFloors := make(chan int)
	//drvObstr := make(chan bool)
	drvStop := make(chan bool)

	elevTx := make(chan elevator.ElevatorMessage)
	elevRx := make(chan elevator.ElevatorMessage)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollStopButton(drvStop)
	go bcast.Receiver(15648, elevRx)
	go bcast.Transmitter(15648, elevTx)

	for {
		select {
		case atFloor := <-drvFloors:
			elevator.FsmFloor(atFloor)
			msg := elevator.ElevatorMessage{ElevatorID,"FLOOR", atFloor, atFloor}
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			fmt.Printf("New Floor Sent\n")
		case stop := <-drvStop:
			elevator.FsmStop(stop)
		case messageReceived := <-elevRx:
			elevator.FsmMessageReceivedHandler(messageReceived,ElevatorID)
			fmt.Printf("New ButtonPress Sent\n")
		case buttonCall := <-drvButtons:
			msg := elevator.ElevatorMessage{ElevatorID,"ORDER", int(buttonCall.Button), buttonCall.Floor}
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			elevTx <- msg
			fmt.Printf("New message sent\n")
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
