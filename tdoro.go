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
	Work float64
	Rest float64
}

func Scheduler(work, rest float64) Schedule {
	return Schedule{
		Work: work,
		Rest: rest,
	}
}

func FormatHalfSeconds(c float64) (string, error) {
	totalSeconds := int(c / 2)
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	timeStr := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

	return timeStr, nil
}

func Run(args []string, stdout, stderr io.Writer, bell ringer) error {
	// a predifined schedule
	var schedule Schedule
	//update schedule on args
	if len(args) >= 3 {
		work, errWork := strconv.ParseFloat(args[1], 64)
		rest, errRest := strconv.ParseFloat(args[2], 64)
		if errWork != nil || errRest != nil {
			stderr.Write([]byte(FAILED_SCHED))
			return errors.New(FAILED_SCHED)
		}
		schedule = Scheduler(float64(work), float64(rest))
	} else {
		schedule = Schedule{
			Work: 0.0001,
			Rest: 0.0001,
		}
	}
	//********PROTOTYPE**********
	// Starts timer
	tickCh := make(chan struct {
		t         time.Time
		count     float64
		clocktype string
	}, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// *schedule, *wg,
	go func() {
		// this combines schedule to half seconds
		current := schedule.Work * 60 * 2
		//ticks every half second
		ticker := time.NewTicker(time.Second / 2)
		tickCh <- struct {
			t         time.Time
			count     float64
			clocktype string
		}{
			t:         time.Time{},
			count:     current,
			clocktype: "work",
		}

		for i := range ticker.C {
			fmt.Printf("work: %v\n", i)
			tickCh <- struct {
				t         time.Time
				count     float64
				clocktype string
			}{
				t:         i,
				count:     current,
				clocktype: "work",
			}
			current -= 1
			if current <= 0 {
				break
			}
		}
		err := RingAlarm(bell)
		if err != nil {
			stderr.Write([]byte(FAILED_BELL))
		}
		current = schedule.Rest * 60 * 2
		for i := range ticker.C {
			fmt.Printf("rest: %v\n", i)
			tickCh <- struct {
				t         time.Time
				count     float64
				clocktype string
			}{
				t:         i,
				count:     current,
				clocktype: "rest",
			}
			current -= 1
			if current <= 0 {
				break
			}
		}
		close(tickCh)
		wg.Done()
	}()

	workArt := art.NewClockHeadWorkReel()
	restArt := art.NewClockHeadRestReel()
	// Draws to terminal every second as a time or progress bar
	for v := range tickCh {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		if v.clocktype == "work" {
			workArt.Draw(stdout)
			workArt.Next()
			fmtCount, err := FormatHalfSeconds(v.count)
			if err != nil {
				fmt.Println("failed to format count")
			}
			fmt.Printf("\t%v\n", fmtCount)
		} else {
			restArt.Draw(stdout)
			restArt.Next()
			fmtCount, err := FormatHalfSeconds(v.count)
			if err != nil {
				fmt.Println("failed to format count")
			}
			fmt.Printf("\t%v\n", fmtCount)
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
