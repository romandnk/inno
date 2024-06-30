package worker

import "time"

type Worker struct {
	Retries int
	Num     int
	Tick    time.Duration
	in      chan any
}

func New(cfg Config, in chan any) *Worker {
	return &Worker{
		Retries: cfg.Retries,
		Num:     cfg.Num,
		Tick:    cfg.Tick,
		in:      in,
	}
}

func (w *Worker) Run() {

}
