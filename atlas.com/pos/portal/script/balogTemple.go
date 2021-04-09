package script

import (
	"log"
)

type balogTemple struct {
}

func BalogTemple() PortalScript {
	return balogTemple{}
}

func (a balogTemple) Name() string {
	return "balogTemple"
}

func (a balogTemple) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(105100000, 2)
	return true
}
