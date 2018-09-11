package main

type Timer interface {
	Tick()
}

type timer struct {
}

func CreateTimer() Timer {
	return &timer{}
}

func (t *timer) Tick() {
}
