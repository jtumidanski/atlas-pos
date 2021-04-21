package script

import (
	"github.com/sirupsen/logrus"
)

type ariantCastle struct {
}

func AriantCastle() PortalScript {
	return ariantCastle{}
}

func (a ariantCastle) Name() string {
	return "ariant_castle"
}

func (a ariantCastle) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.HasItem(4031582) {
		p.PlayPortalSound()
		p.WarpById(260000301, 5)
		return true
	} else {
		p.SendPinkNotice("ENTRY_PASS_NEEDED")
		return false
	}
}
