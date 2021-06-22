package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GotoCastle struct {
}

func (p GotoCastle) Name() string {
	return "gotocastle"
}

func (p GotoCastle) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(2324) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(106020501, 0)
		return true
	}
	script.SendPinkNotice(l, c)("NEED_THORN_REMOVER")
	return false
}
