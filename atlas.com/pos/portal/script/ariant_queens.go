package script

import (
	"github.com/sirupsen/logrus"
)

type ariantQueens struct {
}

func AriantQueens() PortalScript {
	return ariantQueens{}
}

func (a ariantQueens) Name() string {
	return "ariant_queens"
}

func (a ariantQueens) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.Morphed(221005) {
		return false
	} else {
		p.PlayPortalSound()
		p.WarpById(260000300, 7)
		p.SendPinkNotice("PALACE_INTRUDER")
		return true
	}
}