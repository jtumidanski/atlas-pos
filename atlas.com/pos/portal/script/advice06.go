package script

import (
	"log"
)

type advice06 struct {
}

func Advice06() PortalScript {
	return advice06{}
}

func (a advice06) Name() string {
	return "advice06"
}

func (a advice06) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press the #e#b[Up]#k arrow#n to use the portal \r\\and move to the next map.", 230, 5)
	return true
}
