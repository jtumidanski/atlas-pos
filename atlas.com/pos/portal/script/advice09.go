package script

import (
	"log"
)

type advice09 struct {
}

func Advice09() PortalScript {
	return advice09{}
}

func (a advice09) Name() string {
	return "advice09"
}

func (a advice09) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press #e#b[Down]#k on the arrow key#n and#e#b[Alt]#k#n at the same time to jump downwards.", 450, 6)
	return true
}
