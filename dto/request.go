package dto

type SendRequest struct {
	Amount string `json:"amount"`
	Address string `json:"address"`
}