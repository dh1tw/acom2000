package acom2000

import (
	"testing"
	"time"
)

func TestAcom2000_TurnOn(t *testing.T) {

	a, err := NewAcom2000()
	if err != nil {
		t.Fatal(err)
	}

	if err := a.Open(); err != nil {
		t.Fatal(err)
	}

	if err := a.TurnOn(); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 2)
	a.Close()
}

func TestAcom2000_TurnOff(t *testing.T) {

	a, err := NewAcom2000()
	if err != nil {
		t.Fatal(err)
	}

	if err := a.Open(); err != nil {
		t.Fatal(err)
	}

	exitRead := make(chan struct{})
	go a.readSp(exitRead)

	time.Sleep(time.Second * 3)

	if err := a.FindDevices(); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	if err := a.TurnOff(); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 1)

	a.Close()
}

func TestAcom2000_GetIDs(t *testing.T) {
	a, err := NewAcom2000()
	if err != nil {
		t.Fatal(err)
	}

	if err := a.Open(); err != nil {
		t.Fatal(err)
	}

	exitRead := make(chan struct{})

	go a.readSp(exitRead)

	if err := a.FindDevices(); err != nil {
		t.Fatal(err)
	}

	<-exitRead
}
