package script

import (
	"log"
)

type apq02 struct {
}

func Apq02() PortalScript {
	return apq02{}
}

func (a apq02) Name() string {
	return "apq02"
}

func (a apq02) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(670010302, 0)
	return true
}
