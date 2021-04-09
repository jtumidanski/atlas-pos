package script

import (
	"log"
)

type apq01 struct {
}

func Apq01() PortalScript {
	return apq01{}
}

func (a apq01) Name() string {
	return "apq01"
}

func (a apq01) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(670010301, 0)
	return true
}
