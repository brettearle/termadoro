package alarm

import (
	"errors"
	"os/exec"
	"runtime"
	"sync"
)

type Ringer interface {
	Ring() error
}

type Bell struct{}

func (b *Bell) Ring() error {
	if runtime.GOOS == "darwin" {
		err := exec.Command("say", "Time").Run()
		if err != nil {
			return errors.New("say failed")
		}
	} else {
		err := exec.Command("spd-say", "Time").Run()
		if err != nil {
			return errors.New("spd-say failed")
		}
	}
	return nil
}

func RingAlarm(bell Ringer) error {
	errChan := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := bell.Ring()
		if err != nil {
			errChan <- errors.New("Failed to sound bell")
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	err := <-errChan
	if err != nil {
		return err
	}
	return nil
}
