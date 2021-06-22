package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
	"math/rand"
)

type Party3R4pt1 struct {
}

func (p Party3R4pt1) Name() string {
	return "party3_r4pt1"
}

func (p Party3R4pt1) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	portalId := uint32(2)
	if rand.Float64() * 3 > 1 {
		portalId = 2
	}
	script.WarpById(l, c)(920010600, portalId)
	return true
}
