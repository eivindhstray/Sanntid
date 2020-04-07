package variables
//global variables
const N_FLOORS = 4
const N_BUTTON_TYPES = 3
const DOOROPENTIME = 2

const N_ELEVATORS = 2
const ELEVATOR_STATE = 2

const WATCHDOGINTERVAL = 10

var COMMSALIVE = true

type QueueOrderType int

const(
	NONE	QueueOrderType = 0
	REMOTE  			   = 1
	LOCAL    			   = 2
)

var elevatorList[N_ELEVATORS][ELEVATOR_STATE] int //N x 2 array with info on the 
//elevators
//For 2 elev

//			Floor 	Dir
// Elev 1	  x		 y
// Elev 2	  x		 y
// ...
// Elev N	  x		 y

//From this array it will be easy to easy to determine orders


