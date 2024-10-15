package models

import (
	"github.com/gofrs/uuid/v5"
)

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
