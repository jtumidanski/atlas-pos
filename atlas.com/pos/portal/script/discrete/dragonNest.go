package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DragonNest struct {
}

func (p DragonNest) Name() string {
	return "drnNest"
}

func (p DragonNest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(3706) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(240040612, "out00")
		return true
	}

	if processor.QuestStarted(l, c)(100203) || processor.HasItem(l, c)(4001094) {
		//EventManager em = pi.getEventManager("NineSpirit")
		//if (!em.startInstance(pi.getPlayer())) {
		//	pi.sendPinkNotice("SOMEONE_IN_MAP")
		//	return false
		//} else {
		//	pi.playPortalSound()
		//	return true
		//}
		processor.PlayPortalSound(l, c)
		return true
	}
	processor.SendPinkNotice(l, c)("STRANGE_FORCE_2")
	return false
}
