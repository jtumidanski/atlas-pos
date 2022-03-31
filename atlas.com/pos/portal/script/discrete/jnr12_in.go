package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Jnr12In struct {
}

func (p Jnr12In) Name() string {
	return "jnr12_in"
}

func (p Jnr12In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(926110401, 0)
	return true
}
