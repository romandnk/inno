package main

import (
	"go.uber.org/fx"

	"inno/attestation1/internal/app"
)

func main() {
	fx.New(app.App()).Run()
}
