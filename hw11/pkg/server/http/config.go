package httpserver

import "time"

type Config struct {
	Host            string        `env:"HTTP_SERVER_HOST" env-default:"0.0.0.0"`
	Port            int           `env:"HTTP_SERVER_PORT" env-default:"8080"`
	ReadTimeout     time.Duration `env:"HTTP_SERVER_READ_TIMEOUT" env-default:"3s"`
	WriteTimeout    time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT" env-default:"5s"`
	ShutdownTimeout time.Duration `env:"HTTP_SERVER_SHUTDOWN" env-default:"5s"`
}
