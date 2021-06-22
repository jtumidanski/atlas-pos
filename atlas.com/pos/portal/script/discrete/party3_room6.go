package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Room6 struct {
}

func (p Party3Room6) Name() string {
	return "party3_room6"
}

func (p Party3Room6) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010700, 23)
	return true
}
