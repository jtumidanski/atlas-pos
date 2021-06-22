package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Room1 struct {
}

func (p Party3Room1) Name() string {
	return "party3_room1"
}

func (p Party3Room1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010200, 13)
	return true
}
