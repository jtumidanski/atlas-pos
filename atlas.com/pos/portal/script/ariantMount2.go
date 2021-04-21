package script

import (
	"github.com/sirupsen/logrus"
)

type ariantMount2 struct {
}

func AriantMount2() PortalScript {
	return ariantMount2{}
}

func (a ariantMount2) Name() string {
	return "ariantMount2"
}

func (a ariantMount2) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(980010000, 0)
	return true
}
