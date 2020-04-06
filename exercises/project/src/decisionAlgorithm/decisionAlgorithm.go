// New order and remove order on the following form:
// "NewOrder (1), floor, direction"
// "RomeveOrder (0), floor, direction"

//Assume queue has handled the newOrder and removeOrder such that queue is up to date.

package costAlgorithm

import(
	"../elevator"
)

//Very helpful if the messages are stored in a struct on som format like
//struct newOrder (
//		position int
//		direction int
//)

//Positions 0 through N
//Directions: -1 - down, 0 - idle, 1 - up


/*Case 1: Elevator n in Stop:
			if other elevator also in Stop:
				Choose closest one
  Case 2: Elevator n going up:
			pick up anyone at or above going up

  Case 3: Elevator n going down.
			pick up anyone at or below going down
*/
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
