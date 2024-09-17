package core

import (
	"time"
)

func (cpu *CPU) UpdateTimersIfNeeded(lastTimerUpdate *time.Time) {
	if time.Since(*lastTimerUpdate) >= (time.Second / 60) {
		updateTimers(cpu)
		*lastTimerUpdate = time.Now()
	}
}
