package config

import "github.com/amirhnajafiz/hades/internal/config/cronjobs"

type Config struct {
	CronJobs  []cronjobs.Config `koanf:"cronjobs"`
	Interval  int               `koanf:"interval"`
	Threshold int               `koanf:"threshold"`
}
