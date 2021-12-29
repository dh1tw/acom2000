package acom2000

import (
	"errors"
	"fmt"
	"sync"
	"time"

	serial "go.bug.st/serial.v1"
)

type Acom2000 struct {
	sync.Mutex
	sp                    serial.Port
	spConfig              *serial.Mode
	spPort                string
	serialNumber          int
	serialNumberLeftTube  int
	serialNumberRightTube int
	cpuRevision           int
	coverage              Coverage
}

type Coverage int

const (
	CoverageAmateurBands Coverage = iota
	CoverageWithout2428
	CoverageUnlimited
)

type Motor int

const (
	BandMotor Motor = iota
	LoadMotor
	TuneMotor
)

type MotorPosition int

const (
	Light MotorPosition = iota
	Dark
)

type Event byte

const (
	PowerOffExecuted Event = '0'
	PowerOnExecuted  Event = '1'
	HVOnExecuted     Event = '2'
	OffCoolingTubes  Event = '3'
	OffExecuted      Event = '4'
)

func NewAcom2000() (*Acom2000, error) {

	a := &Acom2000{
		spConfig: &serial.Mode{
			BaudRate: 1200,
			Parity:   serial.NoParity,
			DataBits: 8,
			StopBits: serial.OneStopBit,
		},
		spPort: "/dev/ttyUSB0",
	}

	return a, nil
}

func (a *Acom2000) Open() error {

	sp, err := serial.Open(a.spPort, a.spConfig)
	if err != nil {
		return err
	}
	if err := sp.SetRTS(false); err != nil {
		return err
	}
	if err := sp.SetDTR(false); err != nil {
		return err
	}
	a.sp = sp

	exitRead := make(chan struct{})

	go a.readSp(exitRead)

	return nil
}

// async reading
func (a *Acom2000) readSp(close <-chan struct{}) {

	buf := make([]byte, 0, 100)
	for {
		char := make([]byte, 100)
		n, err := a.sp.Read(char)
		if err != nil {
			return
		}

		if n == 0 {
			continue
		}

		if char[0] == '\x00' || char[0] == '\x0d' {
			fmt.Printf("bytes read: %s (%x)\n", string(buf), buf)
			buf = make([]byte, 0, 100) //clear buf
		} else {
			for i := 0; i < n; i++ {
				buf = append(buf, char[i])
			}
		}
	}
}

// type command struct{
// 	to:
// }

// func (a *Acom2000) getCmd(data []byte) ()

func (a *Acom2000) process(data []byte) error {

	return nil
}

func (a *Acom2000) Close() error {
	if a.sp != nil {
		return a.sp.Close()
	}
	return nil
}

func (a *Acom2000) TurnOn() error {
	if a.sp == nil {
		return errors.New("serial port must be initialized to turn on amplifier")
	}

	if err := a.sp.SetRTS(true); err != nil {
		return err
	}
	if err := a.sp.SetDTR(true); err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	if err := a.sp.SetRTS(false); err != nil {
		return err
	}
	if err := a.sp.SetDTR(false); err != nil {
		return err
	}

	return nil
}

func (a *Acom2000) FindDevices() error {

	_, err := a.sp.Write([]byte{'\xFF', '\x71', 'L', '\x00'})

	return err
}

func (a *Acom2000) TurnOff() error {

	_, err := a.sp.Write([]byte{'\x41', '\x71', '0', '\x00'})

	return err
}

func (a *Acom2000) GetSegment(seg int) error {

	ascii, err := dec2Ascii(seg)
	if err != nil {
		return err
	}

	if len(ascii) == 1 {
		ascii = fmt.Sprintf("0%s", ascii)
	}

	fmt.Printf("segment: %d (%s)\n", seg, ascii)

	msg := []byte{'\x41', '\x71', 'W', '1'}
	msg = append(msg, ascii...)
	msg = append(msg, '\x00')

	_, err = a.sp.Write(msg)
	_, err = a.sp.Write([]byte{'\x41', '\x71', 'F', '\x00'})

	return err
}

func (a *Acom2000) MotorPosition(m Motor) (int, MotorPosition, error) {

	return 0, Light, nil
}

func (a *Acom2000) StepUpMotor(m Motor) error {
	return nil
}

func (a *Acom2000) StepDownMotor(m Motor) error {
	return nil
}

func (a *Acom2000) ClearLastFreq() error {
	return nil
}

func (a *Acom2000) Memory(segment int) error {
	return nil
}

func (a *Acom2000) SetMemory(segment int, data []byte) error {
	return nil
}

func (a *Acom2000) Eeprom(page int) error {
	return nil
}

func (a *Acom2000) SetEeprom(page int, data []byte) error {
	return nil
}

func (a *Acom2000) StartAutoTune() error {
	return nil
}

func (a *Acom2000) SetStandby() error {

	_, err := a.sp.Write([]byte{'\x41', '\x71', 'S', '\x00'})
	return err
}

func (a *Acom2000) SetOperation() error {
	_, err := a.sp.Write([]byte{'\x41', '\x71', 'O', '\x00'})
	return err
}

func (a *Acom2000) SerialNumber() int {
	a.Lock()
	defer a.Unlock()
	return a.serialNumber
}

func (a *Acom2000) SerialNumberLeftTube() int {
	a.Lock()
	defer a.Unlock()
	return a.serialNumberLeftTube
}

func (a *Acom2000) SerialNumberRightTube() int {
	a.Lock()
	defer a.Unlock()
	return a.serialNumberRightTube
}

func (a *Acom2000) CPURevision() int {
	a.Lock()
	defer a.Unlock()
	return a.cpuRevision
}

// W1
func (a *Acom2000) LastMotorSetting(segment int, antenna int) error {
	return nil
}

// W4
func (a *Acom2000) LastAntenna(segment int) error {
	return nil
}

// W5
func (a *Acom2000) SetMotorsToDefault(segment int) error {
	return nil
}
