package domain

type Balance struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

func NewBalance(confirmed string, unconfirmed string) *Balance {
	return &Balance{
		Confirmed:   confirmed,
		Unconfirmed: unconfirmed,
	}
}
