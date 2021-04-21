package script

import (
	"github.com/sirupsen/logrus"
)

type cannonTuto07 struct {
}

func CannonTuto07() PortalScript {
	return cannonTuto07{}
}

func (a cannonTuto07) Name() string {
	return "cannon_tuto_07"
}

func (a cannonTuto07) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.SetDirectionStatus(true)
	p.LockUI()
	p.SpawnNPC(579711, 1096012, -51, -97, 0, true)
	p.SetNPCValue(579711, "summon")
	p.UpdateInfo("fly", "579711")
	p.SendDirectionInfo(3, 0)
	p.SendDirectionInfo(3, 2)
	p.SendDirectionInfo(4, 0)
	return true
}
