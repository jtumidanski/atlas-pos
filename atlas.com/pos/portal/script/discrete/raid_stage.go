package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
)

type RaidStage struct {
}

func (p RaidStage) Name() string {
	return "raid_stage"
}

func (p RaidStage) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) == 0 {
		var nextStage uint32

		if c.MapId()%500 >= 100 {
			nextStage = c.MapId() + 100
		} else {
			nextStage = 970030001 + uint32(math.Floor(float64(c.MapId()-970030100)/500))
		}
		processor.PlayPortalSound(l, c)
		processor.WarpRandom(l, span, c)(nextStage)
		return true
	} else {
		processor.SendLightBlueNotice(l, c)("DEFEAT_ALL_MONSTERS")
		return false
	}
}
