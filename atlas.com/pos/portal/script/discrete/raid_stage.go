package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
	"math"
)

type RaidStage struct {
}

func (p RaidStage) Name() string {
	return "raid_stage"
}

func (p RaidStage) Enter(l logrus.FieldLogger, c script.Context) bool {
	if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) == 0 {
		var nextStage uint32

		if c.MapId()%500 >= 100 {
			nextStage = c.MapId() + 100
		} else {
			nextStage = 970030001 + uint32(math.Floor(float64(c.MapId()-970030100)/500))
		}
		script.PlayPortalSound(l, c)
		script.WarpRandom(l, c)(nextStage)
		return true
	} else {
		script.SendLightBlueNotice(l, c)("DEFEAT_ALL_MONSTERS")
		return false
	}
}
