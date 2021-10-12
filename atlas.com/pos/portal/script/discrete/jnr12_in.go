package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr12In struct {
}

func (p Jnr12In) Name() string {
	return "jnr12_in"
}

func (p Jnr12In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(926110401, 0)
	return true
}
