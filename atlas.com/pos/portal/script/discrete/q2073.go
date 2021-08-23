package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Q2073 struct {
}

func (p Q2073) Name() string {
	return "q2073"
}

func (p Q2073) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestStarted(l, c)(2073) {
		script.SendPinkNotice(l, c)("PRIVATE_PROPERTY")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(900000000, 0)
	return true
}