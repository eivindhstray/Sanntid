package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	s, err := net.ResolveUDPAddr("udp", ":20005")

	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp", s)

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		buffer := make([]byte, 1024)
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("->", string(buffer[0:n-1]), '\n')

		data := []byte("hello from udps")

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err, "error")
			return
		}
		time.Sleep(2 * time.Second)
	}
}

//server IP 10.100.23.147
