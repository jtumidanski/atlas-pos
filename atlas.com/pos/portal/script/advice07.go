package script

import (
	"github.com/sirupsen/logrus"
)

type advice07 struct {
}

func Advice07() PortalScript {
	return advice07{}
}

func (a advice07) Name() string {
	return "advice07"
}

func (a advice07) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("You can view the World Map by pressing the #e#b[W]#k#nkey.", 350, 5)
	return true
}
