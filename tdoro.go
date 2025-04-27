package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	art "github.com/brettearle/termadoro/art"
)

const (
	SUCCESS      = "Termadoro Success\n"
	FAILED_BELL  = "Failed to sound bell\n"
	FAILED_SCHED = "Schedule args not numbers\n"
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
		Work: work,
		Rest: rest,
	}
}

func Run(args []string, stdout, stderr io.Writer, bell ringer) error {
	// a predifined schedule
	var schedule Schedule
	//update schedule on args
	if len(args) >= 3 {
		work, errWork := strconv.Atoi(args[1])
		rest, errRest := strconv.Atoi(args[2])
		if errWork != nil || errRest != nil {
			stderr.Write([]byte(FAILED_SCHED))
			return errors.New(FAILED_SCHED)
		}
		schedule = Scheduler(work, rest)
	} else {
		schedule = Schedule{
			Work: 1,
			Rest: 1,
		}
	}
	//********PROTOTYPE**********
	// Starts timer
	tickCh := make(chan struct {
		t         time.Time
		count     int
		clocktype string
	}, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		current := schedule.Work
		ticker := time.NewTicker(time.Millisecond)
		tickCh <- struct {
			t         time.Time
			count     int
			clocktype string
		}{
			t:     time.Time{},
			count: current,
		}

		for i := range ticker.C {
			fmt.Printf("work: %v\n", i)
			tickCh <- struct {
				t         time.Time
				count     int
				clocktype string
			}{
				t:     i,
				count: current,
			}
			current -= 1
			if current == 0 {
				break
			}
		}
		err := RingAlarm(bell)
		if err != nil {
			stderr.Write([]byte(FAILED_BELL))
		}
		current = schedule.Rest
		for i := range ticker.C {
			fmt.Printf("rest: %v\n", i)
			tickCh <- struct {
				t         time.Time
				count     int
				clocktype string
			}{
				t:     i,
				count: current,
			}
			current -= 1
			if current == 0 {
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
		// TODO: rip this out
		if i.count%4 == 0 {
			fmt.Printf("%v\n%v\n", art.ClockHeadWork1, i.count)
		} else if i.count%3 == 0 {
			fmt.Printf("%v\n%v\n", art.ClockHeadWork2, i.count)
		} else if i.count%2 == 0 {
			fmt.Printf("%v\n%v\n", art.ClockHeadWork3, i.count)
		} else {
			fmt.Printf("%v\n%v\n", art.ClockHeadWork4, i.count)
		}
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
