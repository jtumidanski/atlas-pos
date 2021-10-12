package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterMagiclibrar struct {
}

func (p EnterMagiclibrar) Name() string {
	return "enterMagiclibrar"
}

func (p EnterMagiclibrar) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(20718) {
		//var cml = pi.getEventManager("Cygnus_Magic_Library");
		//cml.setProperty("player", pi.getPlayer().getName());
		//cml.startInstance(pi.getPlayer());
	} else {
		script.WarpById(l, span, c)(101000003, 8)
	}
	script.PlayPortalSound(l, c)
	return true
}
