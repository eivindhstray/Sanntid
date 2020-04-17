package main

import (
	"os"
	"fmt"
	"./Network/network/bcast"
	"time"
	"strconv"
)



type msg struct{
	word string
	id int
	number int
}

type alive struct{
	id int
}
var count int

func main(){
	localid,err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Print("panicyo")
	}
	tx := make(chan msg)
	rx := make(chan msg)
	acktx := make(chan alive)
	ackrx := make(chan alive)
	alivemsgtimer := time.NewTimer(0)
	fmt.Print("id: ", localid, "\n")
	
	timeOut := time.NewTimer(0)
	go bcast.Receiver(15648, rx)
	go bcast.Transmitter(15648,tx)
	go bcast.Receiver(15649,ackrx)
	go bcast.Transmitter(15649,acktx)
	alivemsgtimer.Reset(200*time.Millisecond)
	for{
		select{
			case receivedmsg:= <- rx:

				msgid := receivedmsg.id
				num := receivedmsg.number
				fmt.Print("msgreceived id:",receivedmsg.id,"\n")
				switch msgid{
				case 1:
					if localid != 1{
						count = num
						
					}
				
				default:

				}

			case alivemsg := <- ackrx:
				fmt.Print(alivemsg.id,"is id\n")
				if alivemsg.id == localid{
					timeOut.Reset(1*time.Second)
				}
			
		
			
				
			case <- timeOut.C:
					localid = 1
					
			case <- alivemsgtimer.C:
				msg := alive{localid}
				acktx <- msg
				
		
		}
		if localid == 1{

			count = count+1
			fmt.Print("mastercount",count,"\n")
			msg := msg{"Master",1,count}
			tx<-msg
			
	
		}else{
			fmt.Print("slavecount",count,"\n")
			msg := msg{"Slave",2,0}
			tx <- msg
		}
		time.Sleep(1*time.Second)

	}
}
