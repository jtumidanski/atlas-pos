package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk01 struct {
}

func (p EvanTalk01) Name() string {
	return "evantalk01"
}

func (p EvanTalk01) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo01=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o;mo01=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon01")
	return true
}
