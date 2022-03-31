package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GuyFawkes0Floor struct {
}

func (p GuyFawkes0Floor) Name() string {
	return "guyfawkes0_floor"
}

func (p GuyFawkes0Floor) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.WarpById(l, span, c)(674030000, 0)
	return true
}
