package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantMount struct {
}

func (p AriantMount) Name() string {
	return "ariantMount"
}

func (p AriantMount) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(980010020, 0)
	return true
}
