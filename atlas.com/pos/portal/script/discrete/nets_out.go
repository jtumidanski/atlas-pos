package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type NetsOut struct {
}

func (p NetsOut) Name() string {
	return "nets_out"
}

func (p NetsOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	loc, err := processor.GetSavedLocation(l, span, c)("MIRROR")
	if err != nil {
		l.WithError(err).Warnf("Unable to find location to warp to. Character %d may be stuck.", c.CharacterId())
		return false
	}
	if loc.MapId() == 260020500 {
		processor.WarpById(l, span, c)(loc.MapId(), 3)
	} else {
		processor.WarpRandom(l, span, c)(loc.MapId())
	}
	return true
}
