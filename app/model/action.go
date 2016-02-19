package models

import (
	"time"
)

type Action struct {
	ActionId   int64
	ActionCode int
	CreateDate time.Time
	Sent       bool
}

func (a *Action) Save() int64 {

}
