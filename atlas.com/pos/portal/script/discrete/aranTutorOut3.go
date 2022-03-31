package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorOut3 struct {
}

func (p AranTutorOut3) Name() string {
	return "aranTutorOut3"
}

func (p AranTutorOut3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.TeachSkill(l, c)(20000016, 0, -1, -1)
	processor.TeachSkill(l, c)(20000016, 1, 0, -1)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(914000220, 1)
	return true
}
