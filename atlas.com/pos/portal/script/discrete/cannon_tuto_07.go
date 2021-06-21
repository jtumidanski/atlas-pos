package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CannonTuto07 struct {
}

func (p CannonTuto07) Name() string {
	return "cannon_tuto_07"
}

func (p CannonTuto07) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.SetDirectionStatus(l, c)(true)
	script.LockUI(l, c)
	script.SpawnNPC(l, c)(579711, 1096012, -51, -97, 0, true)
	script.SetNPCValue(l, c)(579711, "summon")
	script.UpdateInfo(l, c)("fly", "579711")
	script.SendDirectionInfo(l, c)(3, 0)
	script.SendDirectionInfo(l, c)(3, 2)
	script.SendDirectionInfo(l, c)(4, 0)
	return true
}
