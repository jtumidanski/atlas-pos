package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartToKerning struct {
}

func (p DepartToKerning) Name() string {
	return "Depart_ToKerning"
}

func (p DepartToKerning) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventManager em = pi.getEventManager("KerningTrain")
	//if (!em.startInstance(pi.getPlayer())) {
	//	pi.sendPinkNotice("WAGON_FULL")
	//	return false
	//}
	script.PlayPortalSound(l, c)
	return true
}
