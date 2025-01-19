package domain

import (
	"time"
)

type UserDO struct {
	ID        int64
	Email     string
	Password  string
	Extra     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
