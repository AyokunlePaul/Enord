package model

type Medication struct {
	Id     string  `json:"_id"`
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
	Code   string  `json:"code"`
	Image  string  `json:"image"`
}

func NewMedication(data map[string]interface{}) Medication {
	return Medication{}
}
