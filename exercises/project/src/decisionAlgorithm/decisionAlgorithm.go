// New order and remove order on the following form:
// "NewOrder (1), floor, direction"
// "RomeveOrder (0), floor, direction"

//Assume queue has handled the newOrder and removeOrder such that queue is up to date.

package decisionAlgorithm

import (
	"../elevator"
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
func decisionAlgorithm(newFloor int, CurrentDirection elevator.ElevDir) {
	switch CurrentDirection {
	case Up:
		for i := newFloor; i < variables.N_FLOORS; i++ {
			if elevator.QueueLocal[i][HallUp] == variables.REMOTE {
				elevator.localQueueSetLocal(i, int(elevator.HallUp))
			}
		}
	case Down:
		for i := newFloor - 1; i > -1; i-- {
			if elevator.QueueLocal[i][HallDown] == variables.REMOTE {
				queueLocal[i][HallDown] = variables.LOCAL
			}
		}
	case Stop:
		for i := 0; i < variables.N_FLOORS; i++ {
			if queueLocal[i][HallDown] == variables.REMOTE {
				queueLocal[i][HallDown] = variables.LOCAL
			} else if queueLocal[i][HallUp] == variables.REMOTE {
				queueLocal[i][HallUp] = variables.LOCAL
			}

		}
	}
}

//Returns the floor in the new order
func getNewOrderPosition(elev elevator.Elevator) int {
	return elev.CurrentFloor
}

//Returns the direction to travel in the new order
func getNewOrderDirection(elev elevator.Elevator) int {
	return elev.Dir
}

func costFunction() bool {
	cost := true
	//No cost if the direction of elevator and order is the same and the elevator will
	//evetually reach floor of order while going in set direction

	//No cost if the elevator is not in use

	//High cost if elevator direction and order direction are opposites.
	if elevator.elevatorGetDir() == 1 && getNewOrderDirection() == -1 {
		cost = false
	}

	if elevator.elevatorGetDir() == -1 && getNewOrderDirection() == 1 {
		cost = false
	}

	return cost
}

//Decision to handle new order.
func costDecision() bool {
	x := costFunction()
	if x == true {
		return true
	}
	return false
}
