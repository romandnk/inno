package config

import (
	"fmt"
	"reflect"

	"github.com/ilyakaznacheev/cleanenv"
)

func ParseEnv(cfg any) error {
	if reflect.ValueOf(cfg).Kind() != reflect.Ptr {
		return fmt.Errorf("cfg must be a pointer to a struct")
	}
	return cleanenv.ReadEnv(cfg)
}
