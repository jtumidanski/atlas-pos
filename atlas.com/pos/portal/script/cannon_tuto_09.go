package script

import (
	"github.com/sirupsen/logrus"
)

type cannonTuto09 struct {
}

func CannonTuto09() PortalScript {
	return cannonTuto09{}
}

func (a cannonTuto09) Name() string {
	return "cannon_tuto_09"
}

func (a cannonTuto09) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.OpenNPCWithScript(8, "1096005")
	return true
}