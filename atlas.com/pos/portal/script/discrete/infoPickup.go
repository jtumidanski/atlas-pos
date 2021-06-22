package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type InfoPickup struct {
}

func (p InfoPickup) Name() string {
	return "infoPickup"
}

func (p InfoPickup) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(1035) {
		script.ShowInfo(l, c)("UI/tutorial.img/21")
	}
	script.BlockPortal(l, c)
	return true
}
