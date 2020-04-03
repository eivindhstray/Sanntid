//Network module for elevators
package network

import (
	"time"
)

const (
	peerPort        = "20004"
	getOrderPort    = "30000"
	removeOrderPort = "30001"
	backupPort      = "30002"
	broadcastTime   = 100 * time.Millisecond
)

/*func networkInit() {
	sendNewOrder := make(chan string)
	sendRemoveOrder := make(chan string)

	recieveNewOrder := make(chan string)
	recieveRemoveOrder := make(chan string)

}*/
