package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk21 struct {
}

func (p EvanTalk21) Name() string {
	return "evantalk21"
}

func (p EvanTalk21) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo21=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o;mo11=o;mo20=o;mo21=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon21")
	return true
}