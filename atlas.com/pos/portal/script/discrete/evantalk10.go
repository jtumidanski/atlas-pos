package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk10 struct {
}

func (p EvanTalk10) Name() string {
	return "evantalk10"
}

func (p EvanTalk10) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo10=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o;mo01=o;mo10=0;mo02=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon10")
	return true
}
