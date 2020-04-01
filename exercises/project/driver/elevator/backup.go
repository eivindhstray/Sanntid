//Backup for queue and orders
//Queue will be updated for eventy new order added and every completed order.

//The main program will fetch the backup queue and sync it against the actual queue in case
//of shutdown as backwards recovery

package elevator

import (
	"fmt"

	"../variables"
)

var backupQueue [variables.N_FLOORS][variables.N_BUTTON_TYPES]bool

func backupInit() {
	for i := 0; i < variables.N_FLOORS; i++ {
		for j := 0; j < variables.N_FLOORS; j++ {
			backupQueue[i][j] = queue[i][j]
		}
	}
	fmt.Println("Backup queue initialized")
}

//To be called every time a order is recieved or completed
func backupSync() {
	for i := 0; i < variables.N_FLOORS; i++ {
		for j := 0; j < variables.N_BUTTON_TYPES; j++ {
			if backupQueue[i][j] != queue[i][j] {
				backupQueue[i][j] = queue[i][j]
			}
		}
	}
	fmt.Println("Backup queue synced")
}

//Only to be used in case of backwards recovery
func fetchBackupQueue() {
	for i := 0; i < variables.N_FLOORS; i++ {
		for j := 0; j < variables.N_BUTTON_TYPES; j++ {
			queue[i][j] = backupQueue[i][j]
		}
	}
	fmt.Println("Backup fetched as a result of backwards recovery")
}
