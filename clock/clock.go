package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	m := minute % 60
	h := hour

	h = h + (minute / 60)
	h = h % 24
	if m < 0 {
		h = h - 1
		m = 60 + m
	}
	if h < 0 {
		h = 24 + h
	}
	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute + minutes)
}
