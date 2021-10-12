package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantMount2 struct {
}

func (p AriantMount2) Name() string {
	return "ariantMount2"
}

func (p AriantMount2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(980010000, 0)
	return true
}
