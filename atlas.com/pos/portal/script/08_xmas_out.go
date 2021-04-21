package script

import (
	"github.com/sirupsen/logrus"
)

type christmasOut2008 struct {
}

func ChristmasOut2008() PortalScript {
	return christmasOut2008{}
}

func (c christmasOut2008) Name() string {
	return "08_xmas_out"
}

func (c christmasOut2008) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(context.MapId()-2, 0)
	return true
}