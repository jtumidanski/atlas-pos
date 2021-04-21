package script

import (
	"github.com/sirupsen/logrus"
)

type davyNext1 struct {
}

func DavyNext1() PortalScript {
	return davyNext1{}
}

func (a davyNext1) Name() string {
	return "davy_next1"
}

func (a davyNext1) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.GetEventProperty("stage2") == "3" {
		p.PlayPortalSound()
		p.WarpById(925100200, 0)
		return true
	}

	p.SendPinkNotice("PORTAL_NOT_YET_OPENED")
	return false
}
