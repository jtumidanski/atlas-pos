package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DojangExit struct {
}

func (p DojangExit) Name() string {
	return "dojang_exit"
}

func (p DojangExit) Enter(l logrus.FieldLogger, c script.Context) bool {
	mapId, portalId := script.GetSavedLocation(l, c)("MIRROR")
	if mapId == 0 {
		mapId = 100000000
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(mapId, portalId)
	return true
}
