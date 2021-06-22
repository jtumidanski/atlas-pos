package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Room2 struct {
}

func (p Party3Room2) Name() string {
	return "party3_room2"
}

func (p Party3Room2) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010300, 1)
	return true
}
