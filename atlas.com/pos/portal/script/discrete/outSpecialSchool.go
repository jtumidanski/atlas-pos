package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutSpecialSchool struct {
}

func (p OutSpecialSchool) Name() string {
	return "outSpecialSchool"
}

func (p OutSpecialSchool) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(925040000, 1)
	return true
}
