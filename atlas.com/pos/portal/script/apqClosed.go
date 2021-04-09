package script

import (
	"log"
)

type apqClosed struct {
}

func ApqClosed() PortalScript {
	return apqClosed{}
}

func (a apqClosed) Name() string {
	return "apqClosed"
}

func (a apqClosed) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.SendPinkNotice("GATE_IS_NOT_YET_OPENED")
	return false
}
