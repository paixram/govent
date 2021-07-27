package govent

import (
	"context"
	"fmt"
	"sync"
	"time"

	logger "github.com/liuxsys/govent/libs/logger"
)

/*
       0 = OFF
	   1 = ON
*/

type STATUS int8

var (
	CERO STATUS = 0
	ONE  STATUS = 1
)

type Recive struct {
	Name   string
	Status STATUS
}

type Maker func(resume Recive)

type Funcionality struct {
	f Maker
}

type Event struct {
	Name      string
	State     chan STATUS
	DefaultSt STATUS
}

type EventBox struct {
	m          sync.Mutex
	Name       string
	EventsFunc map[Event]Funcionality
	Log        logger.LogInfo
	Timeout    int16
	//event      chan Event
}

func (e *EventBox) TimeOut(Seconds int16) {
	e.Timeout = Seconds
}

func (e *EventBox) Regis(Event Event, function Maker) error {

	//function(Event)
	//fmt.Println(Event)
	//fmt.Println("No llego")
	e.EventsFunc[Event] = Funcionality{
		f: function,
	}
	return nil
}

func (e *EventBox) Logger() {
	e.Log.Logging()
}

func (e *EventBox) On(Event Event, Wait context.Context) {
	e.m.Lock()
	timeout, _ := context.WithTimeout(Wait, ((time.Duration)(e.Timeout) * time.Second))
	select {
	case changer := <-Event.State:
		if changer != Event.DefaultSt {
			e.Log.Log[fmt.Sprintf("%s status has been changed [%v -> %v]", Event.Name, Event.DefaultSt, changer)] = logger.OK
			r := Recive{Name: Event.Name, Status: changer}
			e.EventsFunc[Event].f(r)
			delete(e.EventsFunc, Event)
			fmt.Println("Evento vaciado")
		} else {
			e.Log.Log[fmt.Sprintf("%s status has not changed or is equal", Event.Name)] = logger.WARN
		}

		//fmt.Println("El estado no ah cambiado")

	case <-timeout.Done():
		e.Log.Log[fmt.Sprintf("%s event time has expired", Event.Name)] = logger.ERROR
		delete(e.EventsFunc, Event)
		fmt.Println("Evento vaciado")
	}
	e.m.Unlock()
}

func (e *EventBox) Change(Event *Event, St STATUS) {

	Event.State <- St

}

type EventBoxIface interface {
	Regis(Event Event, f Maker) error
	On(Event Event, Wait context.Context)
	Change(Event *Event, St STATUS)
	Logger()
	TimeOut(Seconds int16)
}

func NewEventBox(Name string) EventBoxIface {
	e := &EventBox{
		Name:       Name,
		EventsFunc: make(map[Event]Funcionality),
		Log:        logger.LogInfo{Title: Name, Log: make(map[string]logger.LOGINFO)},
		Timeout:    5,
	}
	return e
}

func NewEvent(Name string) Event {
	ev := Event{
		Name:      Name,
		State:     make(chan STATUS, CERO),
		DefaultSt: CERO,
	}

	return ev
}
