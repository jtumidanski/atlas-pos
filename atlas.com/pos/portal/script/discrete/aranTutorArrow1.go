package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorArrow1 struct {
}

func (p AranTutorArrow1) Name() string {
	return "aranTutorArrow1"
}

func (p AranTutorArrow1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "arr1=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;arr1=o;mo1=o;mo2=o;mo3=o;mo4=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
	return true
}
