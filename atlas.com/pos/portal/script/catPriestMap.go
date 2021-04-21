package script

import (
	"github.com/sirupsen/logrus"
)

type catPriestMap struct {
}

func CatPriestMap() PortalScript {
	return catPriestMap{}
}

func (a catPriestMap) Name() string {
	return "catPriest_map"
}

func (a catPriestMap) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(925000000, 2)
	return true
}
