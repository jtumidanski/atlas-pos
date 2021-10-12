package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDCakeEnter struct {
}

func (p MDCakeEnter) Name() string {
	return "MD_cakeEnter"
}

func (p MDCakeEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(674030100,"in00")
	return true
}
