package script

import (
	"log"
)

type catPriestMap struct {
}

func CatPriestMap() PortalScript {
	return catPriestMap{}
}

func (a catPriestMap) Name() string {
	return "catPriest_map"
}

func (a catPriestMap) Enter(l *log.Logger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(925000000, 2)
	return true
}
