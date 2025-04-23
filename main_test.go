package main

import (
	"bytes"
	"errors"
	"testing"
)

type alarmerSuccess struct {
	result string
}

func (ma *alarmerSuccess) Ring() error {
	ma.result = "success"
	return nil
}

type alarmerFail struct {
	result string
}

func (ma *alarmerFail) Ring() error {
	return errors.New("failed")
}

func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Run("Succesful run", func(t *testing.T) {
		stdout := new(bytes.Buffer)
		Run([]string{}, stdout, nil, &alarmerSuccess{})
		want := SUCCESS
		if stdout.String() != want {
			t.Errorf("got %v want %v", stdout.String(), want)
		}
	})

	t.Run("Failed run", func(t *testing.T) {
		stderr := new(bytes.Buffer)
		err := Run(nil, nil, stderr, &alarmerFail{})
		want := FAILED_BELL + FAILED_BELL
		if stderr.String() != want {
			t.Errorf("got %v want %v", stderr.String(), want)
		}
		if err == nil {
			t.Errorf("want error but got nil")
		}

	})

	t.Run("Scheduled Run work arg not int", func(t *testing.T) {
		stderr := new(bytes.Buffer)
		stdout := new(bytes.Buffer)
		Run([]string{"nameOfBinary", "under test", "1"}, stdout, stderr, &alarmerSuccess{})
		want := "Schedule args not numbers\n"
		if stderr.String() != want {
			t.Errorf("got %v want %v", stderr.String(), want)
		}
	})

	t.Run("Scheduled Run rest arg not int", func(t *testing.T) {
		stderr := new(bytes.Buffer)
		stdout := new(bytes.Buffer)
		Run([]string{"nameOfBinary", "1", "under test"}, stdout, stderr, &alarmerSuccess{})
		want := "Schedule args not numbers\n"
		if stderr.String() != want {
			t.Errorf("got %v want %v", stderr.String(), want)
		}
	})
}

func TestScheduler(t *testing.T) {
	t.Run("Scheduler", func(t *testing.T) {
		work := 3
		rest := 2
		got := Scheduler(work, rest)
		if got.Work != work {
			t.Errorf("want %v but got %v", work, got.Work)
		}
		if got.Rest != rest {
			t.Errorf("want %v but got %v", work, got.Rest)
		}
	})
}

func TestAlarm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testAlarm := alarmerSuccess{
			result: "",
		}
		err := RingAlarm(&testAlarm)
		if err != nil {
			t.Errorf("got %v want %v", err, nil)
		}
		if testAlarm.result != "success" {
			t.Errorf("got %v want %v", testAlarm.result, "success")
		}
	})

	t.Run("fail", func(t *testing.T) {
		testAlarm := alarmerFail{
			result: "",
		}
		err := RingAlarm(&testAlarm)
		if err == nil {
			t.Errorf("got %v want error", err)
		}
		if testAlarm.result != "" {
			t.Errorf("got %v want %v", testAlarm.result, "")
		}
	})
}
