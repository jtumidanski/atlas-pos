package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type NetsOut struct {
}

func (p NetsOut) Name() string {
	return "nets_out"
}

func (p NetsOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	mapId, _ := script.GetSavedLocation(l, c)("MIRROR")
	if mapId == 260020500 {
		script.WarpById(l, c)(mapId, 3)
	} else {
		script.WarpRandom(l, c)(mapId)
	}
	return true
}
