package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk11 struct {
}

func (p EvanTalk11) Name() string {
	return "evantalk11"
}

func (p EvanTalk11) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "mo11=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o;mo11=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon11")
	return true
}
