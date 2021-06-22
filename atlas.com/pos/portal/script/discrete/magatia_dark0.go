package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MagatiaDark0 struct {
}

func (p MagatiaDark0) Name() string {
	return "magatia_dark0"
}

func (p MagatiaDark0) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(7770) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(926130000, "out00")
		return true
	}
	script.SendPinkNotice(l, c)("PIPE_TOO_DARK")
	return false
}
