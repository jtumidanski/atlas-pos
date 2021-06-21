package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EvanRoom1 struct {
}

func (p EvanRoom1) Name() string {
	return "evanRoom1"
}

func (p EvanRoom1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "hand=o") {
		return false
	}
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o;mo20=o;hand=o;mo21=o")
	script.ShowInfo(l, c)("UI/tutorial/evan/0/0")
	script.ShowIntro(l, c)("Effect/OnUserEff.img/guideEffect/evanTutorial/evanBalloon70")
	return true
}
