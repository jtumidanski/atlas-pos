package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk00 struct {
}

func (p EvanTalk00) Name() string {
	return "evantalk00"
}

func (p EvanTalk00) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo00=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "mo00=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon00")
	return true
}
