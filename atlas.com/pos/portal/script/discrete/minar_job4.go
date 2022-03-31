package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MinrarJob4 struct {
}

func (p MinrarJob4) Name() string {
	return "minar_job4"
}

func (p MinrarJob4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(240010501, "out00")
	return true
}
