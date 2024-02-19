package config

import (
	"github.com/amirhnajafiz/hades/internal/config/cronjobs"
	"github.com/amirhnajafiz/hades/internal/config/operator"
)

type Config struct {
	Operator  operator.Config   `koanf:"operator"`
	CronJobs  []cronjobs.Config `koanf:"cronjobs"`
	Interval  int               `koanf:"interval"`
	Threshold int               `koanf:"threshold"`
}
