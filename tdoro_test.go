package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	alarm "github.com/brettearle/termadoro/internal"
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
		stderr := new(bytes.Buffer)
		Run([]string{"nameOfBinary", "0.0001", "0.0001"}, stdout, stderr, &alarmerSuccess{})
		// want := SUCCESS
		got := strings.Contains(stdout.String(), SUCCESS)
		if !got {
			t.Errorf("got %v want %v in contains", got, SUCCESS)
		}
	})

	t.Run("Failed run", func(t *testing.T) {
		stderr := new(bytes.Buffer)
		stdout := new(bytes.Buffer)
		err := Run([]string{"nameOfBinary", "0.0001", "0.0001"}, stdout, stderr, &alarmerFail{})
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
		Run([]string{"nameOfBinary", "under test", "0.0001"}, stdout, stderr, &alarmerSuccess{})
		want := "Schedule args not numbers\n"
		if stderr.String() != want {
			t.Errorf("got %v want %v", stderr.String(), want)
		}
	})
	t.Run("Scheduled Run rest arg not int", func(t *testing.T) {
		stderr := new(bytes.Buffer)
		stdout := new(bytes.Buffer)
		Run([]string{"nameOfBinary", "0.0001", "under test"}, stdout, stderr, &alarmerSuccess{})
		want := "Schedule args not numbers\n"
		if stderr.String() != want {
			t.Errorf("got %v want %v", stderr.String(), want)
		}
	})
}

func TestScheduler(t *testing.T) {
	t.Run("Scheduler", func(t *testing.T) {
		work := float64(0.0001)
		rest := float64(0.0001)
		got := Scheduler(work, rest)
		if got.Work != work {
			t.Errorf("want %v but got %v", work, got.Work)
		}
		if got.Rest != rest {
			t.Errorf("want %v but got %v", work, got.Rest)
		}
	})
}

func TestFormatHalfSeconds(t *testing.T) {
	t.Run("FormatCount 6000", func(t *testing.T) {
		count := float64(6000)
		got, _ := FormatHalfSeconds(count)
		want := "00:50:00"
		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("FormatCount 3000", func(t *testing.T) {
		count := float64(3000)
		got, _ := FormatHalfSeconds(count)
		want := "00:25:00"
		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("FormatCount 2678", func(t *testing.T) {
		count := float64(2678)
		got, _ := FormatHalfSeconds(count)
		want := "00:22:19"
		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("FormatCount 2999", func(t *testing.T) {
		count := float64(2999)
		got, _ := FormatHalfSeconds(count)
		want := "00:24:59"
		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}

func TestAlarm(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testAlarm := alarmerSuccess{
			result: "",
		}
		err := alarm.RingAlarm(&testAlarm)
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
		err := alarm.RingAlarm(&testAlarm)
		if err == nil {
			t.Errorf("got %v want error", err)
		}
		if testAlarm.result != "" {
			t.Errorf("got %v want %v", testAlarm.result, "")
		}
	})
}
