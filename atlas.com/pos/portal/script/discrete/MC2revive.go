package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type MC2Revive struct {
}

func (p MC2Revive) Name() string {
	return "MC2revive"
}

func (p MC2Revive) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, c)(c.MapId() - 100)
	return true
}
