package script

import (
	"log"
)

type apq1 struct {
}

func Apq1() PortalScript {
	return apq1{}
}

func (a apq1) Name() string {
	return "apq1"
}

func (a apq1) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(670010400, 0)
	return true
}
