package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DavyNext4 struct {
}

func (p DavyNext4) Name() string {
	return "davy_next4"
}

func (p DavyNext4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	r1, err := processor.ReactorByName(l, span, c)("sMob1")
	if err != nil {
		return false
	}
	r2, err := processor.ReactorByName(l, span, c)("sMob2")
	if err != nil {
		return false
	}
	r3, err := processor.ReactorByName(l, span, c)("sMob3")
	if err != nil {
		return false
	}
	r4, err := processor.ReactorByName(l, span, c)("sMob4")
	if err != nil {
		return false
	}

	if r1.State() >= 1 && r2.State() >= 1 && r3.State() >= 1 && r4.State() >= 1 && processor.MonsterCount(l, c) == 0 {

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

		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(925100500, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
