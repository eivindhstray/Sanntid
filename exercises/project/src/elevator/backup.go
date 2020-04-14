//Backup for queue and orders
//Queue will be updated for eventy new order added and every completed order.

//The main program will fetch the backup queue and sync it against the actual queue in case
//of shutdown as backwards recovery

package elevator

import (
	"fmt"

	"../variables"
)

var backupQueue [variables.N_FLOORS][variables.N_BUTTON_TYPES]variables.QueueOrderType

func BackupSyncQueue() {
	for i := 0; i < variables.N_FLOORS; i++ {
		for j := 0; j < variables.N_FLOORS; j++ {
			if queueLocal[i][j] == variables.LOCAL || queueLocal[i][j] == variables.REMOTE {
				backupQueue[i][j] = variables.REMOTE
			}
		}
	}
	fmt.Println("Backup queue syncreonised")
}

//Only to be used in case of backwards recovery
func fetchBackupQueue() {
	for i := 0; i < variables.N_FLOORS; i++ {
		for j := 0; j < variables.N_BUTTON_TYPES; j++ {
			queueLocal[i][j] = backupQueue[i][j]
		}
	}
	fmt.Println("Backup fetched as a result of backwards recovery")
}
