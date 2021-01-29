package domain

type Host struct {
	Name   string      `json:"name"`
	Host   string      `json:"host"`
	Result *Statistics `json:"result"`
}
