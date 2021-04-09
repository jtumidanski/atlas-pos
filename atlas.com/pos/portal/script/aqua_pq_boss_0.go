package script

import (
	"log"
)

type aquaPQBoss0 struct {
}

func AquaPQBoss0() PortalScript {
	return aquaPQBoss0{}
}

func (a aquaPQBoss0) Name() string {
	return "aqua_pq_boss_0"
}

func (a aquaPQBoss0) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(230040420, 0)
	return true
}
