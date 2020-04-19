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
	Button      int
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

//N x 2 array with floor, direction and connection status bounded to ID.
type ElevatorList [N_ELEVATORS + 1][3]int
