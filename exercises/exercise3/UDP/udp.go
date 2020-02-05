package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	s, err := net.ResolveUDPAddr("udp", ":20128")
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := net.DialUDP("udp", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s \n", c.RemoteAddr().String())

	for {

		_, err := c.Write([]byte("hei"))

		buffer := make([]byte, 1024)
		_, err = c.Write([]byte("udp hello"))
		if err != nil {
			fmt.Println(err)
		}
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("error2", err)
			return
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
		return
		time.Sleep(2 * time.Second)

	}

}
