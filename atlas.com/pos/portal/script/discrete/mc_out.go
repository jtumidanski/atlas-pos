package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type McOut struct {
}

func (p McOut) Name() string {
	return "mc_out"
}

func (p McOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	mapId, portalId := processor.GetSavedLocation(l, c)("MONSTER_CARNIVAL")
	if mapId == 0 {
		mapId = 102000000
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(mapId, portalId)
	return true
}
