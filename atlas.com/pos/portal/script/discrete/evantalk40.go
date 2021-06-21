package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanTalk40 struct {
}

func (p EvanTalk40) Name() string {
	return "evantalk40"
}

func (p EvanTalk40) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22014, "mo40=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22014, "mo30=o;mo40=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon40")
	return true
}
