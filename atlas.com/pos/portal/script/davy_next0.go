package script

import (
	"github.com/sirupsen/logrus"
)

type davyNext0 struct {
}

func DavyNext0() PortalScript {
	return davyNext0{}
}

func (a davyNext0) Name() string {
	return "davy_next0"
}

func (a davyNext0) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.MapMonsterCount(context.MapId()) == 0 && a.passedGrindMode(context.MapId()) {
		p.PlayPortalSound()
		p.WarpById(925100100, 0)
		return true
	} else {
		p.SendPinkNotice("PORTAL_NOT_YET_OPENED")
		return false
	}
}

func (a davyNext0) passedGrindMode(mapId uint32) bool {
	//if (eim.getIntProperty("grindMode") == 0) {
	//	return true
	//}
	//return eim.activatedAllReactorsOnMap(map, 2511000, 2517999)
	return false
}
