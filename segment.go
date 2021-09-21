package main

import (
	"time"
)

type Segment struct {
	Start      time.Duration
	End        time.Duration
	Moment     time.Time
	MomentText string
	GroupName  string
}
