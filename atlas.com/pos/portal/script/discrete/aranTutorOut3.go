package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorOut3 struct {
}

func (p AranTutorOut3) Name() string {
	return "aranTutorOut3"
}

func (p AranTutorOut3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.TeachSkill(l, c)(20000016, 0, -1, -1)
	script.TeachSkill(l, c)(20000016, 1, 0, -1)
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(914000220, 1)
	return true
}
