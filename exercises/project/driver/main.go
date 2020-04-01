package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"./elevator"
	"./elevio"
	"./network/bcast"
	"./variables"
	"./network/peers"
	"./network/localip"
)

func main() {

	elevio.Init("localhost:15658", variables.N_FLOORS)
	

	elevator.ElevatorInit()
	elevator.QueueInit()
	elevator.backupInit()
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
	// We can disable/enable the transmitter after it has been started.
	// This could be used to signal that we are somehow "unavailable".
	peerTxEnable := make(chan bool)
	go peers.Transmitter(15647, id, peerTxEnable)
	go peers.Receiver(15647, peerUpdateCh)

	// Channels
	drvButtons := make(chan elevio.ButtonEvent)
	drvFloors := make(chan int)
	drvObstr := make(chan bool)
	drvStop := make(chan bool)

	elevTx := make(chan elevator.ElevatorMessage)
	elevRx := make(chan elevator.ElevatorMessage)

	go elevio.PollButtons(drvButtons)
	go elevio.PollFloorSensor(drvFloors)
	go elevio.PollObstructionSwitch(drvObstr)
	go elevio.PollStopButton(drvStop)
	go elevator.FsmPollButtonRequest(drvButtons)
	go bcast.Receiver(15647,elevRx)
	go bcast.Transmitter(15648,elevTx)


	for {
		select {
		case a := <-drvFloors:
			elevator.FsmFloor(a)
		case a := <-drvStop:
			elevator.FsmStop(a)
		case p := <-elevRx:
			elevator.FsmMessageReceived(p)
		case s := <-drvButtons:
			msg := elevator.ElevatorMessage{"ORDER", int(s.Button), s.Floor}
			elevTx <- msg
			time.Sleep(1*time.Second)
			fmt.Printf("New message:\n")
		

		//case s:= <- drvFloors:
		//msg := elevator.ElevatorMessage{"FINISHED", s,0}

		case f := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", f.Peers)
			fmt.Printf("  New:      %q\n", f.New)
			fmt.Printf("  Lost:     %q\n", f.Lost)
		}
	}

}

// chmod +x ElevatorServer

// cant just run main. correct command:
// go run *.go
