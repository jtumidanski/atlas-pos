package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MCRevive4 struct {
}

func (p MCRevive4) Name() string {
	return "MCRevive4"
}

func (p MCRevive4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	portalId := uint32(0)
	//switch (pi.getTeam()) {
	//case 0:
	//	portal = 4
	//	break
	//case 1:
	//	portal = 3
	//	break
	//}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(980000401, portalId)
	return true
}
