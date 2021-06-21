package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DavyNext3 struct {
}

func (p DavyNext3) Name() string {
	return "davy_next3"
}

func (p DavyNext3) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.MonsterCount(l, c) == 0 && p.passedGrindMode(c.MapId()) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(925100400, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}

func (p DavyNext3) passedGrindMode(mapId uint32) bool {
	//if (eim.getIntProperty("grindMode") == 0) {
	//	return true
	//}
	//return eim.activatedAllReactorsOnMap(map, 2511000, 2517999)
	return false
}
