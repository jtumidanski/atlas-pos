package script

import (
	"github.com/sirupsen/logrus"
)

type contactDragon struct {
}

func ContactDragon() PortalScript {
	return contactDragon{}
}

func (a contactDragon) Name() string {
	return "contactDragon"
}

func (a contactDragon) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(900090100, 0)
	return true
}
