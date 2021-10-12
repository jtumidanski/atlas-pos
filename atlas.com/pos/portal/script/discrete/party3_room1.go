package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3Room1 struct {
}

func (p Party3Room1) Name() string {
	return "party3_room1"
}

func (p Party3Room1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(920010200, 13)
	return true
}
