package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

type LoginReply struct {
	REplyMssg string
}