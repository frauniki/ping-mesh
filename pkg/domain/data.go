package domain

import "time"

type PushData struct {
	Timestamp time.Time `json:"timestamp"`
	Hosts     []*Host   `json:"hosts"`
}
