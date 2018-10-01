package check

import "time"

type Scheduler struct {
	ToFire      func()
	UpdateCycle time.Duration
	ticker      *time.Ticker
}

func (scheduler Scheduler) Start() {
	scheduler.ticker = time.NewTicker(scheduler.UpdateCycle)
	go func() {
		for {
			select {
			case <-scheduler.ticker.C:
				scheduler.ToFire()
			}
		}
	}()
}

func (scheduler Scheduler) Stop() {
	scheduler.ticker.Stop()
}
