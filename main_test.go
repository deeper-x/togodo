package main

import (
	"testing"
	"time"

	"github.com/deeper-x/togodo/booking"
	"github.com/deeper-x/togodo/communicator"
	"github.com/deeper-x/togodo/mtime"
	"github.com/deeper-x/togodo/settings"
)

func TestEventIsNow(t *testing.T) {
	now := time.Now().Format(settings.UTM)
	eventTS, err := mtime.NowMinute()
	if err != nil {
		t.Error(err)
	}

	res := mtime.EventIsNow(eventTS)
	if err != nil {
		t.Error(err)
	}

	if !res {
		t.Errorf("event %v should be %v", eventTS, now)
	}
}

func TestEventIsNowNot(t *testing.T) {
	now := time.Now().Unix()
	eventTSLate := now + 10

	res := mtime.EventIsNow(eventTSLate)

	if res {
		t.Errorf("event %v should be %v", eventTSLate, now)
	}
}

func TestSendEmail(t *testing.T) {
	_, err := communicator.SendEmail("albertodeprezzo@gmail.com", "hello, todo...")

	if err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	booking.SelectEvents()
}
