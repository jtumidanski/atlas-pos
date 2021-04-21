package script

import (
	"github.com/sirupsen/logrus"
)

type cannonTuto06 struct {
}

func CannonTuto06() PortalScript {
	return cannonTuto06{}
}

func (a cannonTuto06) Name() string {
	return "cannon_tuto_06"
}

func (a cannonTuto06) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.SetDirectionStatus(true)
	p.LockUI()
	p.OpenNPCWithScript(3, "1096003")
	return true
}
