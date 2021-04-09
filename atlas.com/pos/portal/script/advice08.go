package script

import (
	"log"
)

type advice08 struct {
}

func Advice08() PortalScript {
	return advice08{}
}

func (a advice08) Name() string {
	return "advice08"
}

func (a advice08) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("You can check your character's stats by pressing the #e#b[S]#k#nkey.", 350, 5)
	return true
}
