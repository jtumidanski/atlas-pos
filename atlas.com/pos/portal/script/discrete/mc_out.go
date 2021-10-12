package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type McOut struct {
}

func (p McOut) Name() string {
	return "mc_out"
}

func (p McOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	mapId, portalId := script.GetSavedLocation(l, c)("MONSTER_CARNIVAL")
	if mapId == 0 {
		mapId = 102000000
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(mapId, portalId)
	return true
}
