package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanRoom0 struct {
}

func (p EvanRoom0) Name() string {
	return "evanRoom0"
}

func (p EvanRoom0) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22014, "mo30=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22014, "mo30=o")
	script.ShowInfo(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon30")
	return true
}