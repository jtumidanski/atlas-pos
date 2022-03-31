package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DojangExit struct {
}

func (p DojangExit) Name() string {
	return "dojang_exit"
}

func (p DojangExit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	mapId, portalId := processor.GetSavedLocation(l, c)("MIRROR")
	if mapId == 0 {
		mapId = 100000000
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(mapId, portalId)
	return true
}
