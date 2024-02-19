package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"

	"github.com/amirhnajafiz/hades/internal/config/cronjobs"
	"github.com/amirhnajafiz/hades/internal/config/logger"
	"github.com/amirhnajafiz/hades/internal/config/operator"
)

// Config stores the initial values of
// our operator
type Config struct {
	Operator  operator.Config   `koanf:"operator"`
	Logger    logger.Config     `koanf:"logger"`
	CronJobs  []cronjobs.Config `koanf:"cronjobs"`
	Interval  int               `koanf:"interval"`
	Threshold int               `koanf:"threshold"`
}

// New loads config from a yaml file
func New(path string) Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Fatalf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}
