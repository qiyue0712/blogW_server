package models

import "time"

//model

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

type RemoveRequest struct {
	IDList []uint `json:"IDList"`
}
