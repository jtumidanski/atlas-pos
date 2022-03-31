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
	mapId, _ := processor.GetSavedLocation(l, c)("MIRROR")
	if mapId == 260020500 {
		processor.WarpById(l, span, c)(mapId, 3)
	} else {
		processor.WarpRandom(l, span, c)(mapId)
	}
	return true
}
