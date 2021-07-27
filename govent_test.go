package govent

import (
	"context"
	"fmt"
	"testing"
)

func TestGoEvent(t *testing.T) {

	// Eventos
	evento := NewEvent("hello")
	sing := NewEvent("sing")

	// EventoBox
	govent := NewEventBox("Server")
	govent.TimeOut(24)

	govent.Regis(evento, func(resume Recive) {
		fmt.Println("Tester:", resume.Name, resume.Status)

		fmt.Println("hello Event")
	})

	govent.Regis(sing, func(resume Recive) {
		fmt.Println("Sing event was executed")

		fmt.Println("Data: ", resume.Name, resume.Status)
	})

	go govent.On(evento, context.Background())

	go govent.On(sing, context.Background())

	govent.Change(&sing, ONE)

	govent.Change(&evento, ONE)

	govent.Logger()

}
