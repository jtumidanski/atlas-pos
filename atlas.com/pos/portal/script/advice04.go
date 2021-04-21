package script

import (
	"github.com/sirupsen/logrus"
)

type advice04 struct {
}

func Advice04() PortalScript {
	return advice04{}
}

func (a advice04) Name() string {
	return "advice04"
}

func (a advice04) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Click \r\\#b<Sera>", 100, 5)
	return true
}
