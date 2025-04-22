package models

import "time"

type Result struct {
	Server       string
	Port         string
	Status       string
	ResponseTime time.Duration
	Timestamp    time.Time
}
