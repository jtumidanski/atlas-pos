package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Apq02 struct {
}

func (p Apq02) Name() string {
	return "apq02"
}

func (p Apq02) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(670010302, 0)
	return true
}
