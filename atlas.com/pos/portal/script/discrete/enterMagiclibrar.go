package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterMagiclibrar struct {
}

func (p EnterMagiclibrar) Name() string {
	return "enterMagiclibrar"
}

func (p EnterMagiclibrar) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(20718) {
		//var cml = pi.getEventManager("Cygnus_Magic_Library");
		//cml.setProperty("player", pi.getPlayer().getName());
		//cml.startInstance(pi.getPlayer());
	} else {
		processor.WarpById(l, span, c)(101000003, 8)
	}
	processor.PlayPortalSound(l, c)
	return true
}
