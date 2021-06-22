package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MirTalk01 struct {
}

func (p MirTalk01) Name() string {
	return "mirtalk01"
}

func (p MirTalk01) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "dt01=o") {
		return false
	}
	script.MapEffect(l, c)("evan/dragonTalk01")
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o")
	return true
}
