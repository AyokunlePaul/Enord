package model

type Drone struct {
	Id           string  `json:"_id"`
	SerialNumber string  `json:"serial_number"`
	Model        string  `json:"model"`
	Weight       float32 `json:"weight"`
	Battery      string  `json:"battery"`
	State        State   `json:"state"`
}

type State int

const (
	IDLE = iota
	LOADING
	LOADED
	DELIVERING
	DELIVERED
	RETURNING
)

func NewDrone(data map[string]interface{}) Drone {
	return Drone{}
}
