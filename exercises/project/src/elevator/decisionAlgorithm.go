// New order and remove order on the following form:
// "NewOrder (1), floor, direction"
// "RomeveOrder (0), floor, direction"

//Assume queue has handled the newOrder and removeOrder such that queue is up to date.

package elevator

import (
	"../variables"
)

/*Case 1: Elevator n in Stop:
			if other elevator also in Stop:
				Choose closest one
  Case 2: Elevator n going up:
			pick up anyone at or above going up

  Case 3: Elevator n going down.
			pick up anyone at or below going down
*/

//Decision based on current floor and direction of elevator
//New orders have been stored as remote in localQueue. Decision determines if its supposed to be local
func decisionAlgorithm(newFloor int, CurrentDirection ElevDir) {
	switch CurrentDirection {
	case Up:
		for i := newFloor; i < variables.N_FLOORS; i++ {
			//HallUp
			if queueLocal[i][HallUp] == variables.REMOTE {
				localQueueSetLocal(i, int(HallUp))
			}
		}
	case Down:
		for i := newFloor - 1; i > -1; i-- {
			if queueLocal[i][HallDown] == variables.REMOTE {
				localQueueSetLocal(i, int(HallDown))
			}
		}
	//Prioritizing down
	case Stop:
		for i := 0; i < variables.N_FLOORS; i++ {
			if queueLocal[i][HallDown] == variables.REMOTE {
				localQueueSetLocal(i, int(HallDown))
				break
			} else if queueLocal[i][HallUp] == variables.REMOTE {
				localQueueSetLocal(i, int(HallUp))
				break
			}

		}
	}
}
