package script

import (
	"log"
)

type ariantMount2 struct {
}

func AriantMount2() PortalScript {
	return ariantMount2{}
}

func (a ariantMount2) Name() string {
	return "ariantMount2"
}

func (a ariantMount2) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(980010000, 0)
	return true
}
