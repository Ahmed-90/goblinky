package gpio

import (
	"fmt"
	"strings"
	"time"
	"os/exec"
	"strconv"
)

type Led struct {
	Pin      string
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
	if l.step++; l.step > len(l.pattern)-1 {
		l.step = 0
		l.state = false
	}
}

func (l *Led) exec() {
	l.state = !l.state
	s := "0"
	if l.state == true {
		s = "1"
	}

	fmt.Println(s, l)

	exec.Command("gpio", "-g", "write", l.Pin, s).Run()
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
			exec.Command("gpio", "-g", "write", l.Pin, "0").Run()
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



func Make(pin string, dur string, pat string) Led {
	l := Led{}
	l.Pin = pin

	exec.Command("gpio", "-g", "mode", l.Pin, "out").Run()

	l.setDuration(dur)
	l.setPattern(pat)
	l.quit = make(chan struct{})

	return l
}

