package script

import (
	"github.com/sirupsen/logrus"
)

type advice03 struct {
}

func Advice03() PortalScript {
	return advice03{}
}

func (a advice03) Name() string {
	return "advice03"
}

func (a advice03) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press #e#b[Up]#k on the arrow key#n to climb up the ladder or rope.", 350, 5)
	return true
}
