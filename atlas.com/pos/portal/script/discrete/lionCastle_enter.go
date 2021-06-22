package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type LionCastleEnter struct {
}

func (p LionCastleEnter) Name() string {
	return "lionCastle_enter"
}

func (p LionCastleEnter) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(211060010, "west00")
	return true
}
