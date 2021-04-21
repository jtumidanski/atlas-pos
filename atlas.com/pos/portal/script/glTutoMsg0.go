package script

import (
	"github.com/sirupsen/logrus"
)

type glTutoMsg0 struct {
}

func GlTutoMsg0() PortalScript {
	return glTutoMsg0{}
}

func (c glTutoMsg0) Name() string {
	return "glTutoMsg0"
}

func (c glTutoMsg0) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Once you leave this area you won't be able to return.", 150, 5)
	return true
}
