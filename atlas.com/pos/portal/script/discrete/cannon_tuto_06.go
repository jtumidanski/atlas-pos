package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CannonTuto06 struct {
}

func (p CannonTuto06) Name() string {
	return "cannon_tuto_06"
}

func (p CannonTuto06) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.SetDirectionStatus(l, c)(true)
	character.LockUI(l, c)
	script.OpenNPCWithScript(l, c)(3, "1096003")
	return true
}
