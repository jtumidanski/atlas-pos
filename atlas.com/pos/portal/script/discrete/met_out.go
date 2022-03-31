package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MetOut struct {
}

func (p MetOut) Name() string {
	return "met_out"
}

func (p MetOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	mapId, portalId := processor.GetSavedLocation(l, c)("MIRROR")
	if mapId == 0 {
		mapId = 102040000
		portalId = 12
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(mapId, portalId)
	return true
}
