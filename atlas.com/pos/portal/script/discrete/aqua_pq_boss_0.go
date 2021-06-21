package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type AquaPQBoss0 struct {
}

func (p AquaPQBoss0) Name() string {
	return "aqua_pq_boss_0"
}

func (p AquaPQBoss0) Enter(l logrus.FieldLogger, c script.Context) bool {
	character.PlayPortalSound(l)
	character.WarpById(l, c)(230040420, 0)
	return true
}
