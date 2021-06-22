package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type RaidOut struct {
}

func (p RaidOut) Name() string {
	return "raidout"
}

func (p RaidOut) Enter(l logrus.FieldLogger, c script.Context) bool {
	mapId, _ := script.GetSavedLocation(l, c)("BOSS_PQ")
	if mapId == 0 {
		mapId = 100000000
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(mapId, 0)
	return true
}
