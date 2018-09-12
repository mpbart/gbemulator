package main

const (
	DIVIDER_REGISTER uint16 = 0xFF04
	TIMER_REGISTER   uint16 = 0xFF05
)

type Timer interface {
	Tick()
}

type timer struct {
	divider int
	mmu     MMU
}

func CreateTimer(mmu MMU) Timer {
	return &timer{
		divider: 0,
		mmu:     mmu,
	}
}

func (t *timer) Tick() {
	t.divider += 1

	if t.divider == 256 {
		t.incrementDivider()
		t.divider = 0
	}
}

func (t *timer) incrementDivider() {
	t.mmu.WriteByte(DIVIDER_REGISTER, t.mmu.ReadAt(DIVIDER_REGISTER)+1)
}

func (t *timer) incrementCounter() {
	t.mmu.WriteByte(TIMER_REGISTER, t.mmu.ReadAt(TIMER_REGISTER)+1)
}
