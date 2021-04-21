package script

import (
	"github.com/sirupsen/logrus"
)

type apqDoor struct {
}

func ApqDoor() PortalScript {
	return apqDoor{}
}

func (a apqDoor) Name() string {
	return "apqDoor"
}

func (a apqDoor) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)

	name := context.Portal().Name()
	r, err := p.GetReactor(context.MapId(), "gate"+name)
	if err != nil || r.State() != 4 {
		p.SendPinkNotice("GATE_NOT_YET_OPEN")
		return false
	}

	p.PlayPortalSound()
	p.WarpByName(670010600, "gt"+name+"PIB")
	return true
}
