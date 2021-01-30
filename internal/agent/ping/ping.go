package ping

import (
	"time"

	"github.com/frauniki/ping-mesh/pkg/common"
	"github.com/frauniki/ping-mesh/pkg/config"
	"github.com/frauniki/ping-mesh/pkg/domain"
	"github.com/go-ping/ping"
	log "github.com/sirupsen/logrus"
)

const (
	timeoutSec = 10
)

func RunPings(hosts []*domain.Host) error {
	//wg := sync.WaitGroup{}

	for _, host := range hosts {
		{
			pinger, err := ping.NewPinger(host.Host)
			if err != nil {
				log.Error(err)
			}
			pinger.SetPrivileged(true)
			pinger.Count = config.Config.Ping.Count
			if pinger.Interval, err = time.ParseDuration(config.Config.Ping.Interval); err != nil {
				log.Error(err)
			}
			if pinger.Timeout, err = time.ParseDuration(config.Config.Ping.Timeout); err != nil {
				log.Error(err)
			}

			log.WithFields(log.Fields{
				"name": host.Name,
				"host": host.Host,
			}).Debugf("Ping to %s", host.Name)
			if err := pinger.Run(); err != nil {
				log.Error(err)
			}

			var result domain.Statistics
			common.CopyStatistics(pinger.Statistics(), &result)
			host.Result = &result
		}

		// Does not work
		// https://github.com/go-ping/ping/issues/54
		// https://github.com/go-ping/ping/pull/84
		/*
			wg.Add(1)

				go func(host *domain.Host) error {
					defer wg.Done()

					pinger, err := ping.NewPinger(host.Host)
					if err != nil {
						return err
					}
					pinger.SetPrivileged(true)
					pinger.Count = config.Config.Ping.Count
					if pinger.Interval, err = time.ParseDuration(config.Config.Ping.Interval); err != nil {
						return err
					}
					if pinger.Timeout, err = time.ParseDuration(config.Config.Ping.Timeout); err != nil {
						return err
					}

					log.WithFields(log.Fields{
						"name": host.Name,
						"host": host.Host,
					}).Debugf("Ping to %s", host.Name)
					if err := pinger.Run(); err != nil {
						return err
					}

					var result domain.Statistics
					common.CopyStatistics(pinger.Statistics(), &result)
					host.Result = &result

					return nil
				}(host)
		*/
	}
	//wg.Wait()
	return nil
}
