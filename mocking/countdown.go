package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// In test we will send to bytes.Buffer so our tests can capture what data is being generated.
// Slow tests ruin developer productivity.
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {	
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
