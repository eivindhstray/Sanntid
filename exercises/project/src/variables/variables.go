package variables

//global variables
const N_FLOORS = 8
const N_BUTTON_TYPES = 3
const N_ELEVATORS = 2

const DOOROPENTIME = 2
const FAULT_TIME = 7
const ELEV_OFFLINE = 100

var ElevatorID string

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
type ElevatorList [N_ELEVATORS + 1][2]int
