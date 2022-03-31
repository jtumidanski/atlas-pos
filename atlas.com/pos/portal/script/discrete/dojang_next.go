package discrete

import (
	"atlas-pos/character"
	_map "atlas-pos/map"
	"atlas-pos/portal"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

type DojangNext struct {
}

func (p DojangNext) Name() string {
	return "dojang_next"
}

func (p DojangNext) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	current := time.Now().UnixNano()
	if current-processor.NPCCooldown(l, c) < (3000 * time.Millisecond).Milliseconds() {
		return false
	}
	processor.SetNPCCooldown(l, c)(current)
	g, err := processor.ReactorByName(l, span, c)("door")
	if err != nil {
		return false
	}

	if g.State() != 1 && processor.MonsterCount(l, c) != 0 {
		processor.SendPinkNotice(l, c)("DOJO_DOOR_NOT_OPEN")
		return false
	}

	if uint32(math.Floor(float64(c.MapId())/100))%100 >= 38 {
		character.PlayPortalSound(l)
		processor.WarpById(l, span, c)(925020003, 0)
		processor.GainExperience(l, c)(int32(2000 * processor.DojoPoints(l, c)))
		return true
	}

	if ((uint32(math.Floor(float64(c.MapId()+100)/100)) % 100) % 6) != 0 {
		originalPoints := processor.DojoPoints(l, c)
		points := _map.DojoPoints(c.MapId())
		processor.AwardDojoPoints(l, c)(points)
		totalPoints := processor.DojoPoints(l, c)

		processor.SendPinkNotice(l, c)(fmt.Sprintf("You received %d training points. Your total training points score is now %d.", totalPoints-originalPoints, totalPoints))
		character.PlayPortalSound(l)
		processor.WarpById(l, span, c)(c.MapId()+100, 0)
		return true
	}

	if uint32(math.Floor(float64(c.MapId()/10000))) != 92503 {
		originalPoints := processor.DojoPoints(l, c)
		points := _map.DojoPoints(c.MapId())
		processor.AwardDojoPoints(l, c)(points)
		totalPoints := processor.DojoPoints(l, c)

		processor.SendPinkNotice(l, c)(fmt.Sprintf("You received %d training points. Your total training points score is now %d.", totalPoints-originalPoints, totalPoints))
		character.PlayPortalSound(l)
		processor.WarpById(l, span, c)(c.MapId()+100, 0)
		return true
	}

	restMapId := c.MapId() + 100
	for i := int32(0); i < 5; i++ {
		cids := _map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), c.MapId()-100*uint32(i))
		for _, cid := range cids {
			for j := i; j >= 0; j-- {
				originalPoints := processor.DojoPointsOtherCharacter(l, c)(cid)
				points := _map.DojoPoints(c.MapId() - 100*uint32(j))
				processor.AwardDojoPointsOtherCharacter(l, c)(cid, points)
				totalPoints := processor.DojoPointsOtherCharacter(l, c)(cid)
				character.SendPinkNotice(l)(c.WorldId(), c.ChannelId(), cid, fmt.Sprintf("You received %d training points. Your total training points score is now %d.", totalPoints-originalPoints, totalPoints))
			}
			portal.WarpById(l, span)(c.WorldId(), c.ChannelId(), cid, restMapId, 0)
		}
	}
	return true
}
