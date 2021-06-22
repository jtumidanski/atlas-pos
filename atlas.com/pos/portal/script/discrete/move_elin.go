package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MoveElin struct {
}

func (p MoveElin) Name() string {
	return "move_elin"
}

func (p MoveElin) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(300000100, "out00")
	script.SendPinkNotice(l, c)("PASSING_TIME_GATE")
	return true
}
