package script

import (
	"log"
)

type adviceMap struct {
}

func AdviceMap() PortalScript {
	return adviceMap{}
}

func (a adviceMap) Name() string {
	return "adviceMap"
}

func (a adviceMap) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press the #e#b[Up]#k arrow#n to use the portal and move to the next map.", 230, 5)
	return true
}
