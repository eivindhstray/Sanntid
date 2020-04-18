package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"./network/bcast"
)

//the message won't send numbers for some reason.
//therefore, we have left the message as it is now, 
//believing our solution should work nicely once the message problem is fixed

var message msg

type msg struct {
	msg    string
	id     int
	number int
}

var status alive

type alive struct {
	id int
}

var count int

func main() {
	localid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print("panicyo")
	}
	message.id = int(localid)
	message.msg = "Hello"
	status.id = int(localid)
	tx := make(chan msg)
	rx := make(chan msg)
	acktx := make(chan alive)
	ackrx := make(chan alive)
	alivemsgtimer := time.NewTimer(0)
	sendtimer := time.NewTimer(0)
	fmt.Print("id: ", localid, "\n")
	timeOut := time.NewTimer(0)
	go bcast.Receiver(15647, rx)
	go bcast.Transmitter(15647, tx)
	go bcast.Receiver(15647, ackrx)
	go bcast.Transmitter(15647, acktx)
	sendtimer.Reset(3*time.Second)
	for {
		select {
		case receivedmsg := <-rx:

			msgid := receivedmsg.id
			fmt.Println("idmsg: ", msgid)
			fmt.Print("\n",receivedmsg.number,"number")
			fmt.Print("\n yoyo", receivedmsg.msg)
			num := receivedmsg.number
			if msgid == 1 && localid == 2 {
					message.number = num
			}else if msgid != localid{
				alivemsgtimer.Reset(1 * time.Second)
			}

		case alivemsg := <-ackrx:
			fmt.Print("msgid\n",alivemsg.id)
			
			timeOut.Reset(1* time.Second)
			

		case <-timeOut.C:
			localid = 1

		case <-alivemsgtimer.C:
			acktx <- alive{localid}

		case<-sendtimer.C:
			if localid == 1 {
			count = count + 1
			message.number = count
				tx <- message
			}
			if	localid == 2 {
				tx <- message
			sendtimer.Reset(3*time.Second)
		}

		
		}
	}
}
