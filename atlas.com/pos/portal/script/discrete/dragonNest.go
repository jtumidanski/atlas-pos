package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DragonNest struct {
}

func (p DragonNest) Name() string {
	return "drnNest"
}

func (p DragonNest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestCompleted(l, c)(3706) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(240040612, "out00")
		return true
	}

	if script.QuestStarted(l, c)(100203) || script.HasItem(l, c)(4001094) {
		//EventManager em = pi.getEventManager("NineSpirit")
		//if (!em.startInstance(pi.getPlayer())) {
		//	pi.sendPinkNotice("SOMEONE_IN_MAP")
		//	return false
		//} else {
		//	pi.playPortalSound()
		//	return true
		//}
		script.PlayPortalSound(l, c)
		return true
	}
	script.SendPinkNotice(l, c)("STRANGE_FORCE_2")
	return false
}
