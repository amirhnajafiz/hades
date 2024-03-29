package config

import (
	"github.com/amirhnajafiz/hades/internal/config/logger"
	"github.com/amirhnajafiz/hades/internal/config/operator"
)

// Default values for the app configs
func Default() Config {
	return Config{
		Operator:  operator.Config{},
		Logger:    logger.Config{},
		CronJobs:  []string{},
		Interval:  0,
		Threshold: 0,
	}
}
