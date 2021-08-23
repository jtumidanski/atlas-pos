package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Room3 struct {
}

func (p Party3Room3) Name() string {
	return "party3_room3"
}

func (p Party3Room3) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010400, 8)
	return true
}