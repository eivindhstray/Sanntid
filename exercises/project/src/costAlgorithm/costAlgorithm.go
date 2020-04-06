// New order and remove order on the following form:
// "NewOrder (1), floor, direction"
// "RomeveOrder (0), floor, direction"

//Assume queue has handled the newOrder and removeOrder such that queue is up to date.

package decision

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




