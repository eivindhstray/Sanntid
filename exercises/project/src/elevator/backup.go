//Backup for queue and orders
//Queue will be updated for eventy new order added and every completed order.

//The main program will fetch the backup queue and sync it against the actual queue in case
//of shutdown as backwards recovery

package elevator

import (
	"fmt"

	"../variables"
)

var BackUpQueue [variables.N_FLOORS][variables.N_BUTTON_TYPES]variables.QueueOrderType

func BackUpQueueInit(){
	BackUpQueue = queue
}

//Backup queue containing every order; local or remote. Not for cab calls, since this should only contain active orders in a functional elevator.
func BackupSyncQueue() {
	for floor := 0; floor < variables.N_FLOORS; floor++ {
		for buttons := 0; buttons < variables.N_BUTTON_TYPES-1; buttons++ {
			if queue[floor][buttons] == variables.LOCAL || queue[floor][buttons] == variables.REMOTE {
				BackUpQueue[floor][buttons] = variables.REMOTE
			}
		}
	}
	fmt.Println("Backup queue synced")
}

//Making local queue identical to backup queue
//Only to be used in case of backwards recovery
func fetchBackupQueue() {
	for i := 0; i < variables.N_FLOORS; i++ {
		for j := 0; j < variables.N_BUTTON_TYPES; j++ {
			queue[i][j] = BackUpQueue[i][j]
		}
	}
	fmt.Println("Backup fetched as a result of backwards recovery")
}

func GetBackUpQueue() [variables.N_FLOORS][variables.N_BUTTON_TYPES]variables.QueueOrderType {
	fetchBackupQueue()
	return BackUpQueue
}
