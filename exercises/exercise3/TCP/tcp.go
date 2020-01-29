package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var count = 0

func handleConnection(c net.Conn) {
	fmt.Print(".")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"

		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide port number")
		return
	}

	Port := ":" + arguments[1]
	l, err := net.Listen("tcp4", Port)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		s, err := net.ResolveTCPAddr("tcp4", Port)
		connection, err := net.DialTCP("tcp4", nil, s)
		connection.Write([]byte(string("yolo")))
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(connection)
		go handleConnection(c)
		count++
		time.Sleep(2 * time.Second)
	}

}
