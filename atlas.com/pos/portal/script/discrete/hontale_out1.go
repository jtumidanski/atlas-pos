package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HorntaleOut1 struct {
}

func (p HorntaleOut1) Name() string {
	return "hontale_out1"
}

func (p HorntaleOut1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(240050400, "sp")
	return true
}
