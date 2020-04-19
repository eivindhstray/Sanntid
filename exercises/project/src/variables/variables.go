package variables

//global variables
const N_FLOORS = 4
const N_BUTTON_TYPES = 3
const DOOROPENTIME = 2

const N_ELEVATORS = 2
const ELEVATOR_STATE = 2
const FAULT_TIME = 7
const NEW_FLOOR_TIMEOUT_PENALTY = 100
const ELEV_OFFLINE = 100

var ElevatorID string

//message

type ElevatorMessage struct {
	ElevID      int
	MessageType string
	Button      int //0 = hallup 1= halldown 2 = cab
	Floor       int
	Dir         int
	Elevators   ElevatorList
}

type QueueOrderType int

const (
	NONE   QueueOrderType = 0
	REMOTE                = 1
	LOCAL                 = 2
)

type ElevatorList [N_ELEVATORS + 1][3]int //N x 2 array with info on the
//elevators
//For 2 elev

//			Floor 	Connection Status
// Elev 1	  x		 	y
// Elev 2	  x		 	y
// ...
// Elev N	  x		 	y

//From this array it will be easy to easy to determine orders

type CabCalls [N_ELEVATORS][4]int

// 			Floor 0		Floor 1		Floor 2		Floor 3
//Elev 1
//Elev 2
//...
//Elev N
