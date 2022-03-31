package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type PPinkOut struct {
}

func (p PPinkOut) Name() string {
	return "PPinkOut"
}

func (p PPinkOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpRandom(l, span, c)(270050000)
	return true
}
