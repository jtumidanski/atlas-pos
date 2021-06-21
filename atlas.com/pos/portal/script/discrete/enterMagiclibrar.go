package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type EnterMagiclibrar struct {
}

func (p EnterMagiclibrar) Name() string {
	return "enterMagiclibrar"
}

func (p EnterMagiclibrar) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestStarted(l, c)(20718) {
		//var cml = pi.getEventManager("Cygnus_Magic_Library");
		//cml.setProperty("player", pi.getPlayer().getName());
		//cml.startInstance(pi.getPlayer());
	} else {
		character.WarpById(l, c)(101000003, 8)
	}
	character.PlayPortalSound(l)
	return true
}
