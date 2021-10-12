package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DavyNext4 struct {
}

func (p DavyNext4) Name() string {
	return "davy_next4"
}

func (p DavyNext4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.ReactorByName(l, c)("sMob1").State() >= 1 &&
		script.ReactorByName(l, c)("sMob2").State() >= 1 &&
		script.ReactorByName(l, c)("sMob3").State() >= 1 &&
		script.ReactorByName(l, c)("sMob4").State() >= 1 &&
		script.MonsterCount(l, c) == 0 {

		//if (eim.getProperty("spawnedBoss") == null) {
		//	int level = (eim.getProperty("level")).toInteger()
		//	int chests = (eim.getProperty("openedChests")).toInteger()
		//
		//	Optional<MapleMonster> boss
		//	if (chests == 0) {
		//		//lord pirate
		//		boss = MapleLifeFactory.getMonster(9300119)
		//	} else if (chests == 1) {
		//		//angry lord pirate
		//		boss = MapleLifeFactory.getMonster(9300105)
		//	} else {
		//		//enraged lord pirate
		//		boss = MapleLifeFactory.getMonster(9300106)
		//	}
		//
		//	boss.ifPresent({ monster ->
		//		monster.changeDifficulty(level, true)
		//		pi.getMap(925100500).spawnMonsterOnGroundBelow(monster, new Point(777, 140))
		//		eim.setProperty("spawnedBoss", "true")
		//	})
		//}

		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(925100500, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
