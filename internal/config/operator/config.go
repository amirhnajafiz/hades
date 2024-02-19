package operator

type Config struct {
	Metrics     int  `koanf:"metrics"`
	Probe       int  `koanf:"prob"`
	LeaderElect bool `koanf:"leader_elect"`
}
