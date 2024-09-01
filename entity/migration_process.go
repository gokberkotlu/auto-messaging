package entity

import (
	"time"
)

type MigrationProcessStatusType uint8

const (
	LastStatus       MigrationProcessStatusType = 0
	MessageBatchLoad MigrationProcessStatusType = 0
)

type MigrationFunc func()

type MigrationProcessItem struct {
	Status MigrationProcessStatusType
	Action MigrationFunc
}

type MigrationProcess struct {
	ID        uint64                     `json:"id" gorm:"primaryKey"`
	Status    MigrationProcessStatusType `json:"status" gorm:"not null"`
	CreatedAt time.Time                  `json:"created_at" gorm:"type:timestamp without time zone;default:now()"`
	UpdatedAt time.Time                  `json:"updated_at" gorm:"type:timestamp without time zone;default:now()"`
}
