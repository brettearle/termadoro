package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

const (
	SUCCESS     = "Termadoro Success"
	FAILED_BELL = "Failed to sound bell"
)

type ringer interface {
	Ring() error
}

type bell struct{}

func (b *bell) Ring() error {
	err := exec.Command("spd-say", "b b b b").Run()
	if err != nil {
		return errors.New("spd-say failed")
	}
	return nil
}

func RingAlarm(bell ringer) error {
	err := bell.Ring()
	if err != nil {
		return errors.New(FAILED_BELL)
	}
	return nil
}

func Run(stdout io.Writer, bell ringer) error {
	// a predifined config

	// Starts timer

	// Draws to terminal every second as a time or progress bar

	// Pings when done
	err := RingAlarm(bell)
	if err != nil {
		stdout.Write([]byte(FAILED_BELL))
		return errors.New(FAILED_BELL)
	}

	// success
	stdout.Write([]byte(SUCCESS))
	return nil
}

func main() {
	Run(os.Stdout, &bell{})
}
