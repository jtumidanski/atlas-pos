package script

import (
	"github.com/sirupsen/logrus"
)

type ariantMount struct {
}

func AriantMount() PortalScript {
	return ariantMount{}
}

func (a ariantMount) Name() string {
	return "ariantMount"
}

func (a ariantMount) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(980010020, 0)
	return true
}
