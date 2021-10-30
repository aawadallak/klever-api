package domain

type Total struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}

func NewTotal(sent string, received string) *Total {
	return &Total{Sent: sent, Received: received}
}

