package domain

// Ref: https://github.com/go-ping/ping/blob/3300c582a663c5bce1d0ecf1868107da0ae38829/ping.go#L204
type Statistics struct {
	PacketsRecv int        `json:"packet_receive"`
	PacketsSent int        `json:"packet_sent"`
	PacketLoss  float64    `json:"packet_loss"`
	IPAddr      *IPAddr    `json:"ip_address"`
	Addr        string     `json:"address"`
	Rtts        []Duration `json:"rtts"`
	MinRtt      Duration   `json:"min_rtt"`
	MaxRtt      Duration   `json:"max_rtt"`
	AvgRtt      Duration   `json:"avg_rtt"`
	StdDevRtt   Duration   `json:"standard_deviation_rtt"`
}
