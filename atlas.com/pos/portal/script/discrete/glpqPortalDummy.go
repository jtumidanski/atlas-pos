package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"github.com/sirupsen/logrus"
)

type GlpqPortalDummy struct {
}

func (p GlpqPortalDummy) Name() string {
	return "glpqPortalDummy"
}

func (p GlpqPortalDummy) Enter(l logrus.FieldLogger, c script.Context) bool {
	r := reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), "mob0")
	if r.State() >= 1 {
		script.SendPinkNotice(l, c)("PORTAL_MALFUNCTION")
		return false
	}

	reactor.ForceHit(l)(c.WorldId(), c.ChannelId(), c.MapId(), r.Id(), 1)
	//EventInstanceManager eim = pi.getEventInstance()
	//eim.setIntProperty("glpq1", 1)
	script.SendPinkNotice(l, c)("STRANGE_FORCE")
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(610030100, 0)
	//pi.getEventInstance().showClearEffect()
	//eim.giveEventPlayersStageReward(1)
	return true
}
