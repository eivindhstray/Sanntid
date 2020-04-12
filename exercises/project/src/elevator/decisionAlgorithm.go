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
/*
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
}*/

//decisionAlgorithm making a choice every time a new order is stashed into queue remote.
//The algorithm tuns through the queue to find the remote order.
//The elevator with the lowest cost makes the order local.

func decisionAlgorithm() {
	var CostArray [variables.N_ELEVATORS + 1]int
	var correctFloor int
	var correctButton int
	//Init cost array
	for elev := 1; elev < variables.N_ELEVATORS+1; elev++ {
		CostArray[elev] = 0
	}

	//Find the remote order and determine cost for every elevator
	for elevator := 1; elevator < variables.N_ELEVATORS+1; elevator++ {
		cost := 0
		for floors := 0; floors < variables.N_FLOORS; floors++ {
			for buttons := 0; buttons < variables.N_BUTTON_TYPES-1; buttons++ {
				if queueLocal[floors][buttons] == variables.REMOTE {
					cost = floors - elev.ElevState[elevator][0]
					if cost < 0 {
						cost = -cost
					}
					correctFloor = floors
					correctButton = buttons
					if correctFloor > elev.ElevState[elevator][0] && elev.dir == Down {
						cost = cost + 10
					}
					if correctFloor < elev.ElevState[elevator][0] && elev.dir == Up {
						cost = cost + 10
					}
					CostArray[elevator] = cost
				}
			}
		}
	}

	//Find best elevator
	var bestElev int
	bestElev = 1
	for elevator := 2; elevator < variables.N_ELEVATORS+1; elevator++ {
		if CostArray[elevator] < CostArray[bestElev] {
			bestElev = elevator
		}

	}

	//Set local in queue of best elevator
	if bestElev == elev.ElevID {
		queueLocal[correctFloor][correctButton] = variables.LOCAL
	}
}
