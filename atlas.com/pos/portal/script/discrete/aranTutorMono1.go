package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AranTutorMono1 struct {
}

func (p AranTutorMono1) Name() string {
	return "aranTutorMono1"
}

func (p AranTutorMono1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)

	if script.ContainsAreaInfo(l, c)(21002, "mo2=o") {
		return false
	}
	character.PlaySound(l, c)("Aran/balloon")
	script.UpdateAreaInfo(l, c)(21002, "mo1=o;mo2=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/aranTutorial/legendBalloon2")
	return true
}
