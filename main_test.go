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

	t.Run("Succesful run", func(t *testing.T) {
		stdout := new(bytes.Buffer)
		Run(stdout, &alarmerSuccess{})
		want := SUCCESS
		if stdout.String() != want {
			t.Errorf("got %v want %v", stdout.String(), want)
		}
	})

	t.Run("Failed run", func(t *testing.T) {
		stdout := new(bytes.Buffer)
		err := Run(stdout, &alarmerFail{})
		want := FAILED_BELL
		if stdout.String() != want {
			t.Errorf("got %v want %v", stdout.String(), want)
		}
		if err == nil {
			t.Errorf("want error but got nil")
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
