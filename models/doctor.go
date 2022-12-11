package models

type Doctor struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Ratings     string `json:"ratings"`
}

type CreateDoctor struct {
	Name        string `json:"name" binding:"required"`
	Designation string `json:"designation" binding:"required"`
	Ratings     string `json:"ratings" binding:"required"`
}
