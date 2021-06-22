package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Room5 struct {
}

func (p Party3Room5) Name() string {
	return "party3_room5"
}

func (p Party3Room5) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010600, 17)
	return true
}
