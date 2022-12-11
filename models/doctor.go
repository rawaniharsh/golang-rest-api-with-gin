package models

type Doctor struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Ratings     string `json:"ratings"`
}
