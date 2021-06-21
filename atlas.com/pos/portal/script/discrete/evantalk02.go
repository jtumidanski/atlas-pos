package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk02 struct {
}

func (p EvanTalk02) Name() string {
	return "evantalk02"
}

func (p EvanTalk02) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo02=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o;mo01=o;mo02=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon02")
	return true
}
