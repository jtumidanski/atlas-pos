package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DavyNext0 struct {
}

func (p DavyNext0) Name() string {
	return "davy_next0"
}

func (p DavyNext0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.MonsterCount(l, c) == 0 && p.passedGrindMode(c.MapId()) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(925100100, 0)
		return true
	} else {
		processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
		return false
	}
}

func (p DavyNext0) passedGrindMode(mapId uint32) bool {
	//if (eim.getIntProperty("grindMode") == 0) {
	//	return true
	//}
	//return eim.activatedAllReactorsOnMap(map, 2511000, 2517999)
	return false
}
