package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CannonTuto09 struct {
}

func (p CannonTuto09) Name() string {
	return "cannon_tuto_09"
}

func (p CannonTuto09) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.OpenNPCWithScript(l, c)(8, "1096005")
	return true
}
