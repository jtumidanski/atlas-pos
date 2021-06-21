package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorMono0 struct {
}

func (p AranTutorMono0) Name() string {
	return "aranTutorMono0"
}

func (p AranTutorMono0) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "mo1=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "mo1=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon1")
	return true
}
