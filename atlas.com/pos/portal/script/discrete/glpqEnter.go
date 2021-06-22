package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqEnter struct {
}

func (p GlpqEnter) Name() string {
	return "glpqEnter"
}

func (p GlpqEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.HasItem(l, c)(3992041) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(610030020, "out00")
		return true
	}
	script.SendPinkNotice(l, c)("GIANT_GATE_NO_BUDGE")
	return false
}
