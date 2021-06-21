package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DojangUp struct {
}

func (p DojangUp) Name() string {
	return "dojang_up"
}

func (p DojangUp) Enter(l logrus.FieldLogger, c script.Context) bool {
	m := script.MonsterById(l, c)(9300216)
	if m == nil {
		script.SendPinkNotice(l, c)("DOJO_MORE_MONSTERS")
		script.EnableActions(l, c)
		return true
	}

	//pi.goDojoUp()
	//pi.getPlayer().getMap().setReactorState()
	//int stage = Math.floor(pi.getPlayer().getMapId() / 100) % 100
	//if ((stage - (stage / 6) | 0) == pi.getPlayer().getVanquisherStage() && !GameConstants.isDojoPartyArea(pi.getPlayer().getMapId())) // we can also try 5 * stage / 6 | 0 + 1
	//{
	//	pi.getPlayer().setVanquisherKills(pi.getPlayer().getVanquisherKills() + 1)
	//}
	script.EnableActions(l, c)
	return true
}
