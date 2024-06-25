package main

import "fmt"

func main() {
	sl := SlowLogger{}
	sl.Info("error loading data", "slow logger", fmt.Errorf("error connection"))

	fl := FastLogger{}
	fl.Info("error loading data", "fast logger", fmt.Errorf("error connection"))
}
