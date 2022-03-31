package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutSpecialSchool struct {
}

func (p OutSpecialSchool) Name() string {
	return "outSpecialSchool"
}

func (p OutSpecialSchool) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(925040000, 1)
	return true
}
