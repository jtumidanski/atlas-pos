package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk20 struct {
}

func (p EvanTalk20) Name() string {
	return "evantalk20"
}

func (p EvanTalk20) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo20=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o;mo11=o;mo20=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon20")
	return true
}
