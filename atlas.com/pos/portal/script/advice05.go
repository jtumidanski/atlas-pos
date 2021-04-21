package script

import (
	"github.com/sirupsen/logrus"
)

type advice05 struct {
}

func Advice05() PortalScript {
	return advice05{}
}

func (a advice05) Name() string {
	return "advice05"
}

func (a advice05) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press #e#b[Q]#k#n to view the Quest window.", 250, 5)
	return true
}
