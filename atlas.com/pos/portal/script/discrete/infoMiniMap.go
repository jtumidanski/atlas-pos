package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type InfoMiniMap struct {
}

func (p InfoMiniMap) Name() string {
	return "infoMinimap"
}

func (p InfoMiniMap) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(1031) {
		script.ShowInfo(l, c)("UI/tutorial.img/25")
	}
	script.BlockPortal(l, c)
	return true
}
