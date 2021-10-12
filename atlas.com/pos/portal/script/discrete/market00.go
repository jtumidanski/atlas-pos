package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Market00 struct {
}

func (p Market00) Name() string {
	return "market00"
}

func (p Market00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(script.GetMarketPortal(l, span, c))
	return true
}
