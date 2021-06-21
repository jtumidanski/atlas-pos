package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide0 struct {
}

func (p AranTutorGuide0) Name() string {
	return "aranTutorGuide0"
}

func (p AranTutorGuide0) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "normal=o") {
		return false
	}
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide1")
	script.SendPinkNotice(l, c)("ARAN_TUTORIAL_REGULAR_ATTACK")
	script.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;mo1=o;mo2=o;mo3=o")
	return true
}
