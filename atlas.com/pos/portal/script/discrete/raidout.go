package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RaidOut struct {
}

func (p RaidOut) Name() string {
	return "raidout"
}

func (p RaidOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	mapId, _ := processor.GetSavedLocation(l, c)("BOSS_PQ")
	if mapId == 0 {
		mapId = 100000000
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(mapId, 0)
	return true
}
