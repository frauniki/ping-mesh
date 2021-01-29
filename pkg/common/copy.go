package common

import (
	"github.com/frauniki/ping-mesh/pkg/domain"
	"github.com/go-ping/ping"
)

func CopyStatistics(src *ping.Statistics, dst *domain.Statistics) {
	dst.PacketsRecv = src.PacketsRecv
	dst.PacketsSent = src.PacketsSent
	dst.PacketLoss = src.PacketLoss
	dst.IPAddr = &domain.IPAddr{
		IP:   src.IPAddr.IP,
		Zone: src.IPAddr.Zone,
	}
	dst.Addr = src.Addr
	dst.MinRtt = domain.Duration{src.MinRtt}
	dst.MaxRtt = domain.Duration{src.MaxRtt}
	dst.AvgRtt = domain.Duration{src.AvgRtt}
	dst.StdDevRtt = domain.Duration{src.StdDevRtt}

	for _, r := range src.Rtts {
		dst.Rtts = append(dst.Rtts, domain.Duration{r})
	}
}
