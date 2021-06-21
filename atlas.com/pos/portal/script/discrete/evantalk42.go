package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk42 struct {
}

func (p EvanTalk42) Name() string {
	return "evantalk42"
}

func (p EvanTalk42) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22014, "mo42=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22014, "mo30=o;mo40=o;mo41=o;mo42=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon42")
	return true
}
