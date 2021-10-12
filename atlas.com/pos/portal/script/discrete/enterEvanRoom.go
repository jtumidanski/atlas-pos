package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterEvanRoom struct {
}

func (p EnterEvanRoom) Name() string {
	return "enterEvanRoom"
}

func (p EnterEvanRoom) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(100030100, 0)
	return true
}
