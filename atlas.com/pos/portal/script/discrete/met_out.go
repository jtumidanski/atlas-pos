package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MetOut struct {
}

func (p MetOut) Name() string {
	return "met_out"
}

func (p MetOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpToSavedLocationDefaultPortal(l, span, c)("MIRROR", 102040000, 12)
	return true
}
