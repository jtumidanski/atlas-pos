package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DojangTuto struct {
}

func (p DojangTuto) Name() string {
	return "dojang_tuto"
}

func (p DojangTuto) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	m := script.MonsterById(l, c)(9300216)
	if m == nil {
		script.SendPinkNotice(l, c)("DOJO_DONT_RUN")
		return false
	}

	//pi.getPlayer().enteredScript("dojang_Msg", pi.getPlayer().getMap().getId())
	//pi.getPlayer().setFinishedDojoTutorial()
	//pi.getClient().getChannelServer().resetDojo(pi.getPlayer().getMap().getId())
	//pi.getClient().getChannelServer().dismissDojoSchedule(pi.getPlayer().getMap().getId(), pi.getParty().get())
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(925020001, 0)
	return true
}
