package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MirTalk00 struct {
}

func (p MirTalk00) Name() string {
	return "mirtalk00"
}

func (p MirTalk00) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.BlockPortal(l, c)
	if script.ContainsAreaInfo(l, c)(22013, "dt00=o") {
		return false
	}
	script.MapEffect(l, c)("evan/dragonTalk00")
	script.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o")
	return true
}
