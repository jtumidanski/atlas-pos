package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorMono3 struct {
}

func (p AranTutorMono3) Name() string {
	return "aranTutorMono3"
}

func (p AranTutorMono3) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "mo4=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;mo1=o;mo2=o;mo3=o;mo4=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon6")
	return true
}
