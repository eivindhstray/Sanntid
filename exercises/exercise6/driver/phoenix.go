package main

import (
	"encoding/binary"
	. "fmt"
	"log"
	"net"
	"os/exec"
	t "time"
)

var counter uint64
var buf = make([]byte, 16)

func spawnBackup() {
	(exec.Command("gnome terminal", "-x", "sh", "-c", "go run phoenix.go")).Run()

	println("New backup running")
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	isPrimary := false
	conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		log.Println("Error: something went wrong")
	}
	log.Println("Yo mofo im the new guy in town")

	if !isPrimary {
		conn.SetReadDeadLine(t.Now().Add(2 * t.Second))
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			isPrimary = true
		} else {
			counter = binary.BigEndian.Uint64(buf[:n])
		}
	}
	conn.Close()

	println("Addr", addr)
	spawnBackup()
	println("New pimp in town")
	bcastconn, _ := net.DialUDP("udp", nil, addr)

	for {
		if counter%10 == 0 {
			println("\t*--------")
			println("\t| Number: ", counter, "\t|")
		} else {
			println("\t| Number: ", counter, "\t|")
		}
		counter++
		binary.BigEndian.PutUint64(buf, counter)
		_, _ = bcastConn.Write(buf)
		t.sleep(100 * t.Millisecond)
	}

}
