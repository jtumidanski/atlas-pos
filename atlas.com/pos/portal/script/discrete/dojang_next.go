package discrete

import (
	"atlas-pos/character"
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
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
	if current-script.NPCCooldown(l, c) < (3000 * time.Millisecond).Milliseconds() {
		return false
	}
	script.SetNPCCooldown(l, c)(current)
	g := script.ReactorByName(l, c)("door")
	if g == nil {
		return false
	}

	if g.State() != 1 && script.MonsterCount(l, c) != 0 {
		script.SendPinkNotice(l, c)("DOJO_DOOR_NOT_OPEN")
		return false
	}

	if uint32(math.Floor(float64(c.MapId())/100))%100 >= 38 {
		character.PlayPortalSound(l)
		script.WarpById(l, span, c)(925020003, 0)
		script.GainExperience(l, c)(int32(2000 * script.DojoPoints(l, c)))
		return true
	}

	if ((uint32(math.Floor(float64(c.MapId()+100)/100)) % 100) % 6) != 0 {
		originalPoints := script.DojoPoints(l, c)
		points := _map.DojoPoints(c.MapId())
		script.AwardDojoPoints(l, c)(points)
		totalPoints := script.DojoPoints(l, c)

		script.SendPinkNotice(l, c)(fmt.Sprintf("You received %d training points. Your total training points score is now %d.", totalPoints-originalPoints, totalPoints))
		character.PlayPortalSound(l)
		script.WarpById(l, span, c)(c.MapId()+100, 0)
		return true
	}

	if uint32(math.Floor(float64(c.MapId()/10000))) != 92503 {
		originalPoints := script.DojoPoints(l, c)
		points := _map.DojoPoints(c.MapId())
		script.AwardDojoPoints(l, c)(points)
		totalPoints := script.DojoPoints(l, c)

		script.SendPinkNotice(l, c)(fmt.Sprintf("You received %d training points. Your total training points score is now %d.", totalPoints-originalPoints, totalPoints))
		character.PlayPortalSound(l)
		script.WarpById(l, span, c)(c.MapId()+100, 0)
		return true
	}

	restMapId := c.MapId() + 100
	for i := int32(0); i < 5; i++ {
		cids := _map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), c.MapId()-100*uint32(i))
		for _, cid := range cids {
			for j := i; j >= 0; j-- {
				originalPoints := script.DojoPointsOtherCharacter(l, c)(cid)
				points := _map.DojoPoints(c.MapId() - 100*uint32(j))
				script.AwardDojoPointsOtherCharacter(l, c)(cid, points)
				totalPoints := script.DojoPointsOtherCharacter(l, c)(cid)
				character.SendPinkNotice(l)(c.WorldId(), c.ChannelId(), cid, fmt.Sprintf("You received %d training points. Your total training points score is now %d.", totalPoints-originalPoints, totalPoints))
			}
			character.WarpById(l, span)(c.WorldId(), c.ChannelId(), cid, restMapId, 0)
		}
	}
	return true
}
