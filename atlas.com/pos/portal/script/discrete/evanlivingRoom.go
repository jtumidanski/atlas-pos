package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanLivingRoom struct {
}

func (p EvanLivingRoom) Name() string {
	return "evanlivingRoom"
}

func (p EvanLivingRoom) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(100030102, "in00")
	return true
}
