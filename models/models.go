package models

type Employee struct {
	Id      string `json:"id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
