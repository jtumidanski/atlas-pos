package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorArrow3 struct {
}

func (p AranTutorArrow3) Name() string {
	return "aranTutorArrow3"
}

func (p AranTutorArrow3) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "arr3=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(21002, "cmd=o;normal=o;arr0=o;arr1=o;arr2=o;arr3=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialArrow1")
	return true
}
