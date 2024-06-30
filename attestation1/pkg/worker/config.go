package worker

import "time"

type Config struct {
	Num     int           `env:"WORKER_NUM" env-default:"3"`
	Tick    time.Duration `env:"WORKER_TICK" env-default:"1s"`
	Retries int           `env:"WORKER_RETRIES" env-default:"3"`
}
