package worker

import "time"

type Config struct {
	Num     int           `yaml:"num"`
	Tick    time.Duration `yaml:"tick"`
	Retries int           `yaml:"retries"`
}
