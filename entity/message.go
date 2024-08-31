package entity

import (
	"time"
)

type StatusType uint8

const (
	Active StatusType = 0
	Sent   StatusType = 1
)

type Message struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	To        string     `json: "to" gorm:"size:20;not null"`
	Status    StatusType `json: "status" gorm:"not null"`
	Content   string     `json: "content" gorm:"size:200;not null"`
	CreatedAt time.Time  `json: "created_at" gorm:"type:timestamp without time zone;default:now()"`
	UpdatedAt time.Time  `json: "updated_at" gorm:"type:timestamp without time zone;default:now()"`
}
