package script

import (
	"log"
)

type advice01 struct {
}

func Advice01() PortalScript {
	return advice01{}
}

func (a advice01) Name() string {
	return "advice01"
}

func (a advice01) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Click \r\\#b<Heena>#k", 100, 5)
	return true
}
