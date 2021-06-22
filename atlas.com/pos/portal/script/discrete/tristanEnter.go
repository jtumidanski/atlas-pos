package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TristanEnter struct {
}

func (p TristanEnter) Name() string {
	return "tristanEnter"
}

func (p TristanEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !script.QuestCompleted(l, c)(2238) {
		script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(105100101, "in00")
	return true
}
