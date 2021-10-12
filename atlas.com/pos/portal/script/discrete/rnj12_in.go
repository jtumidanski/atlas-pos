package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Rnj12In struct {
}

func (p Rnj12In) Name() string {
	return "rnj12_in"
}

func (p Rnj12In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(926100401, 0)
	return true
}
