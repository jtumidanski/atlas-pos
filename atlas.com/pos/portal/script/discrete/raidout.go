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
	loc, err := processor.GetSavedLocation(l, span, c)("BOSS_PQ")
	if err != nil {
		l.WithError(err).Warnf("Unable to find location to warp to. Character %d may be stuck.", c.CharacterId())
		return false
	}
	processor.PlayPortalSound(l, c)
	if loc.MapId() == 0 {
		processor.WarpById(l, span, c)(100000000, 0)
	} else {
		processor.WarpById(l, span, c)(loc.MapId(), 0)
	}
	return true
}
