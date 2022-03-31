package generic

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func EnterMiniDungeon(l logrus.FieldLogger, span opentracing.Span, c script.Context) func(baseId uint32, dungeonId uint32, count uint8) bool {
	return func(baseId uint32, dungeonId uint32, count uint8) bool {
		if c.MapId() == baseId {
			_, partyExists := processor.GetParty(l, c)
			if partyExists {
				//if (pi.isLeader()) {
				//	for (int i = 0; i < dungeons; i++) {
				//		if (pi.startDungeonInstance(dungeonId + i)) {
				//			pi.playPortalSound()
				//			pi.warp(dungeonId + i, "out00")
				//			return true
				//		}
				//	}
				//} else {
				processor.SendPinkNotice(l, c)("PARTY_LEADER_MUST_ENTER")
				return false
				//}
			} else {
				//for (int i = 0; i < dungeons; i++) {
				//	if (pi.startDungeonInstance(dungeonId + i)) {
				//		pi.playPortalSound()
				//		pi.warp(dungeonId + i, "out00")
				//		return true
				//	}
				//}
			}
			processor.SendPinkNotice(l, c)("ALL_MINI_DUNGEON_IN_USE")
			return false
		} else {
			processor.PlayPortalSound(l, c)
			processor.WarpByName(l, span, c)(baseId, "MD00")
			return true
		}
	}
}
