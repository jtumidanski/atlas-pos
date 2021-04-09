package script

import (
	"log"
)

type apq3 struct {
}

func Apq3() PortalScript {
	return apq3{}
}

func (a apq3) Name() string {
	return "apq3"
}

func (a apq3) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(670010600, 0)
	return true
}
