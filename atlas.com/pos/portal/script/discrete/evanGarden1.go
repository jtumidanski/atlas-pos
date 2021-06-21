package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanGarden1 struct {
}

func (p EvanGarden1) Name() string {
	return "evanGarden1"
}

func (p EvanGarden1) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(22008) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(100030103, "west00")
	} else {
		script.SendPinkNotice(l, c)("CANNOT_ENTER_BACKYARD_WITHOUT")
	}
	return true
}
