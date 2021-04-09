package script

import (
	"log"
)

type balogEnd struct {
}

func BalogEnd() PortalScript {
	return balogEnd{}
}

func (a balogEnd) Name() string {
	return "balog_end"
}

func (a balogEnd) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	if p.CanHold(4001261, 1) {
		p.GainItem(4001261, 1)
		p.PlayPortalSound()
		p.WarpById(105100100, 0)
		return true
	}

	p.SendPinkNotice("INVENTORY_FULL_ERROR")
	return false
}
