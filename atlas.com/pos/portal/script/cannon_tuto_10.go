package script

import (
	"log"
)

type cannonTuto10 struct {
}

func CannonTuto10() PortalScript {
	return cannonTuto10{}
}

func (a cannonTuto10) Name() string {
	return "cannon_tuto_10"
}

func (a cannonTuto10) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.UnlockUI()
	return true
}
