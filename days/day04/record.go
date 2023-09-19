package day04

import "time"

type Action string

type Record struct {
	guardId   int
	action    Action
	timeStamp time.Time
}

const (
	begins  Action = "begins"
	sleeps  Action = "sleeps"
	wakesUp Action = "wakes up"
)
