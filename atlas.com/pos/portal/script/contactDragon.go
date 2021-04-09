package script

import (
	"log"
)

type contactDragon struct {
}

func ContactDragon() PortalScript {
	return contactDragon{}
}

func (a contactDragon) Name() string {
	return "contactDragon"
}

func (a contactDragon) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(900090100, 0)
	return true
}
