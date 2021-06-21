package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorGuide1 struct {
}

func (p AranTutorGuide1) Name() string {
	return "aranTutorGuide1"
}

func (p AranTutorGuide1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "chain=o") {
		return false
	}
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/tutorialGuide2")
	character.SendPinkNotice(l, c)("ARAN_TUTORIAL_CONSECUTIVE")
	script.UpdateAreaInfo(l, c)(21002, "normal=o;arr0=o;arr1=o;mo1=o;chain=o;mo2=o;mo3=o;mo4=o")
	return true
}
