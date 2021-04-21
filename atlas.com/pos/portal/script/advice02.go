package script

import (
	"github.com/sirupsen/logrus"
)

type advice02 struct {
}

func Advice02() PortalScript {
	return advice02{}
}

func (a advice02) Name() string {
	return "advice02"
}

func (a advice02) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press #e#b[Alt]#k#n to\r\\ JUMP.", 100, 5)
	return true
}
