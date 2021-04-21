package script

import (
	"github.com/sirupsen/logrus"
)

type advice00 struct {
}

func Advice00() PortalScript {
	return advice00{}
}

func (a advice00) Name() string {
	return "advice00"
}

func (a advice00) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("You can move by using the arrow keys.", 250, 5)
	return true
}
