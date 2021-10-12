package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanFall struct {
}

func (p EvanFall) Name() string {
	return "evanFall"
}

func (p EvanFall) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(900090102, 0)
	return true
}
