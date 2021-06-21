package discrete

import (
	"atlas-pos/character"
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DavyNext2 struct {
}

func (p DavyNext2) Name() string {
	return "davy_next2"
}

func (p DavyNext2) Enter(l logrus.FieldLogger, c script.Context) bool {
	if _map.MonsterCount(l, c)(c.MapId()) == 0 && p.passedGrindMode(c.MapId()) {
		character.PlayPortalSound(l)
		character.WarpById(l, c)(925100300, 0)
		return true
	} else {
		character.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
		return false
	}
}

func (p DavyNext2) passedGrindMode(mapId uint32) bool {
	//if (eim.getIntProperty("grindMode") == 0) {
	//	return true
	//}
	//return eim.activatedAllReactorsOnMap(map, 2511000, 2517999)
	return false
}
