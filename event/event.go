package event

import (
	"time"
)

type Event struct {
	T   time.Time
	Msg string
}

func NewWTime(t time.Time, msg string) Event {
	return Event{T: t, Msg: msg}
}

func NewWDuration(d time.Duration, msg string) Event {
	return Event{T: time.Now().Add(d), Msg: msg}
}
