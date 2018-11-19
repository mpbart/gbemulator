package main

type Timer interface {
	Tick()
}

type timer struct {
	dividerCounter    int
	timerCounter      int
	currentInputClock int
	mmu               MMU
}

func CreateTimer(mmu MMU) Timer {
	return &timer{
		dividerCounter:    0,
		timerCounter:      0,
		currentInputClock: getInputClock(mmu),
		mmu:               mmu,
	}
}

func (t *timer) Tick() {
	t.dividerCounter += 1
	t.timerCounter += 1

	if t.dividerCounter >= 256 {
		t.incrementDividerRegister()
		t.dividerCounter = 0
	}

	if t.timerEnabled() {
		if clk := getInputClock(t.mmu); clk != t.currentInputClock {
			t.currentInputClock = clk
		}

		if t.timerCounter >= t.currentInputClock {
			t.incrementTimerRegister()
			t.timerCounter = 0
		}
	}
}

func (t *timer) incrementDividerRegister() {
	t.mmu.WriteByte(DIVIDER_REGISTER, t.mmu.ReadAt(DIVIDER_REGISTER)+1)
}

func (t *timer) incrementTimerRegister() {
	currentValue := t.mmu.ReadAt(TIMER_REGISTER)
	if currentValue == 0xFF {
		t.mmu.WriteByte(TIMER_REGISTER, t.mmu.ReadAt(TIMER_MODULO))
		t.mmu.FireInterrupt(TIMER_INTERRUPT)
	} else {
		t.mmu.WriteByte(TIMER_REGISTER, currentValue+1)
	}
}

func (t *timer) timerEnabled() bool {
	return GetBit(t.mmu.ReadAt(TIMER_CONTROL), 2) == 1
}

func getInputClock(mmu MMU) int {
	value := mmu.ReadAt(TIMER_CONTROL)
	switch value {
	case 0:
		return 1024
	case 1:
		return 16
	case 2:
		return 64
	default:
		return 256
	}
}
