package config

import "github.com/spf13/viper"

func setDefaultValues() {
	// Global
	viper.SetDefault("interval", "5s")
	viper.SetDefault("log_level", "info")

	// Global.Ping
	viper.SetDefault("ping.count", 1)
	viper.SetDefault("ping.interval", "1s")
	viper.SetDefault("ping.timeout", "2s")
}
