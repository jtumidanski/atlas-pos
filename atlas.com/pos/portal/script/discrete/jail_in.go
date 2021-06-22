package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type JailIn struct {
}

func (p JailIn) Name() string {
	return "jail_in"
}

func (p JailIn) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, c)(300000012,"portal")
	return true
}
