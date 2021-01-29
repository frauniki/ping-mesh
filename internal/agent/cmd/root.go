package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/frauniki/ping-mesh/internal/agent/client"
	"github.com/frauniki/ping-mesh/internal/agent/ping"
	"github.com/frauniki/ping-mesh/internal/agent/ticker"
	"github.com/frauniki/ping-mesh/pkg/config"
	"github.com/frauniki/ping-mesh/pkg/domain"
	"github.com/frauniki/ping-mesh/pkg/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	defaultConfigPath = "configs/"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Start ping-mesh agent")

		var hosts []*domain.Host
		for _, host := range config.Config.Hosts {
			hosts = append(hosts, &domain.Host{
				Name: host.Name,
				Host: host.Host,
			})
		}

		ticker := ticker.NewTicker()
		for range ticker.C {
			if err := ping.RunPings(hosts); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			if config.Config.ServerURL == "" {
				continue
			}
			if err := client.Post(&domain.PushData{
				Timestamp: time.Now(),
				Hosts:     hosts,
			}); err != nil {
				log.Error(err)
			}
		}
	},
}

func init() {
	cobra.OnInitialize(func() {
		if _, err := config.Load(configFile); err != nil {
			log.Fatal(err)
		}
		if err := logger.Init(); err != nil {
			log.Fatal(err)
		}
	})

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./configs/config.yaml)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
