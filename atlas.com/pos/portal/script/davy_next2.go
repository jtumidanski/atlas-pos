package script

import (
	"github.com/sirupsen/logrus"
)

type davyNext2 struct {
}

func DavyNext2() PortalScript {
	return davyNext2{}
}

func (a davyNext2) Name() string {
	return "davy_next2"
}

func (a davyNext2) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.MapMonsterCount(context.MapId()) == 0 && a.passedGrindMode(context.MapId()) {
		p.PlayPortalSound()
		p.WarpById(925100300, 0)
		return true
	} else {
		p.SendPinkNotice("PORTAL_NOT_YET_OPENED")
		return false
	}
}

func (a davyNext2) passedGrindMode(mapId uint32) bool {
	//if (eim.getIntProperty("grindMode") == 0) {
	//	return true
	//}
	//return eim.activatedAllReactorsOnMap(map, 2511000, 2517999)
	return false
}
