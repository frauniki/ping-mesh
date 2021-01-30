package domain

import "time"

type PushData struct {
	AgentHostname string    `json:"agent_hostname"`
	Timestamp     time.Time `json:"timestamp"`
	Hosts         []*Host   `json:"hosts"`
}
