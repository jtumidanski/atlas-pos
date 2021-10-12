package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3Room5 struct {
}

func (p Party3Room5) Name() string {
	return "party3_room5"
}

func (p Party3Room5) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(920010600, 17)
	return true
}
