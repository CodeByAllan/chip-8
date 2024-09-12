package cpu

import (
	"chip-8/common"
	"chip-8/instructions"
	"time"
)

func UpdateTimersIfNeeded(cpu *common.CPU, lastTimerUpdate *time.Time) {
	if time.Since(*lastTimerUpdate) >= (time.Second / 60) {
		instructions.UpdateTimers(cpu)
		*lastTimerUpdate = time.Now()
	}
}
