package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Party3Room4 struct {
}

func (p Party3Room4) Name() string {
	return "party3_room4"
}

func (p Party3Room4) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(920010500, 3)
	return true
}
