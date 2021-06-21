package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk41 struct {
}

func (p EvanTalk41) Name() string {
	return "evantalk41"
}

func (p EvanTalk41) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22014, "mo41=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22014, "mo30=o;mo40=o;mo41=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon41")
	return true
}
