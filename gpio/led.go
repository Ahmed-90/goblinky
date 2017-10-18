package gpio

import (
	"time"
	"strings"
	"strconv"
	"fmt"
)

type Led struct {
	Pin      int
	pattern  []int
	duration time.Duration
	step     int
	state    bool
	quit     chan struct{}
	stepper  *time.Timer
}

func (l *Led) Blink() {
	l.nextStep()
	go l.runner()
}

func (l *Led) nextStep() {
	l.stepper = time.NewTimer(l.duration * time.Duration(l.pattern[l.step]))
	if l.step++; l.step > len(l.pattern) - 1{
		l.step = 0
	}
}

func (l *Led) exec() {
	fmt.Println(!l.state, l)
	l.state = !l.state
}

func (l *Led) Stop() {
	l.quit <- struct{}{}
}

func (l *Led) runner() {
	defer func() {
		l.step = 0
		l.state = false
	}()

	l.exec()

	for {
		select {
		case <-l.quit:
			fmt.Println("Quitting")
			return
		case <-l.stepper.C:
			l.exec()
			l.nextStep()
		}
	}
}

func (l *Led) setPattern(p string) {
	l.pattern = []int{}
	for _, i := range strings.Split(p, ",") {
		t, _ := strconv.Atoi(i)
		l.pattern = append(l.pattern, t)
	}
}

func (l *Led) setDuration(d string) {
	if d == "ms" {
		l.duration = time.Millisecond
	} else {
		l.duration = time.Second
	}
}

func (l *Led) Update(d string, p string) {
	l.setDuration(d)
	l.setPattern(p)
	l.Stop()
	time.Sleep(time.Second * 1)
	l.Blink()
}

func Make(p int, dur string, pat string) Led {
	l := Led{Pin:p}
	l.setDuration(dur)
	l.setPattern(pat)
	l.quit = make(chan struct{})

	return l
}