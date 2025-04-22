package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

const (
	SUCCESS     = "Termadoro Success\n"
	FAILED_BELL = "Failed to sound bell\n"
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

type Schedule struct {
	Work int
	Rest int
}

func Scheduler(work, rest int) Schedule {
	return Schedule{
		Work: 0,
		Rest: 0,
	}
}

func Run(args []string, stdout, stderr io.Writer, bell ringer) error {
	//********PROTOTYPE**********
	// a predifined schedule
	schedule := struct {
		workInMins int
		restInMins int
	}{
		workInMins: 10,
		restInMins: 2,
	}
	// Starts timer
	tickCh := make(chan struct {
		t     time.Time
		count int
	}, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		capacity := schedule.workInMins
		current := 0
		ticker := time.NewTicker(time.Second)
		for i := range ticker.C {
			fmt.Printf("work: %v\n", i)
			tickCh <- struct {
				t     time.Time
				count int
			}{
				t:     i,
				count: current,
			}
			current += 1
			if current == capacity {
				break
			}
		}
		err := RingAlarm(bell)
		if err != nil {
			stderr.Write([]byte(FAILED_BELL))
		}
		capacity = schedule.restInMins
		current = 0
		for i := range ticker.C {
			fmt.Printf("rest: %v\n", i)
			tickCh <- struct {
				t     time.Time
				count int
			}{
				t:     i,
				count: current,
			}
			current += 1
			if current == capacity {
				break
			}
		}
		close(tickCh)
		wg.Done()
	}()

	// Draws to terminal every second as a time or progress bar
	for i := range tickCh {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		str := "#"
		for idx := i.count; idx > 0; idx-- {
			str = str + "#"
		}
		fmt.Printf("%v\n%v\n", str, i.count)
	}

	wg.Wait()
	fmt.Println("ticker done")
	//********END OF PROTOTYPE*****************

	// Pings when done
	err := RingAlarm(bell)
	if err != nil {
		stderr.Write([]byte(FAILED_BELL))
		return errors.New(FAILED_BELL)
	}

	// success
	stdout.Write([]byte(SUCCESS))
	return nil
}

func main() {
	Run(os.Args, os.Stdout, os.Stderr, &bell{})
}
