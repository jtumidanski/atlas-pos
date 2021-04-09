package script

import (
	"log"
)

type apq00 struct {
}

func Apq00() PortalScript {
	return apq00{}
}

func (a apq00) Name() string {
	return "apq00"
}

func (a apq00) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(670010300, 0)
	return true
}
