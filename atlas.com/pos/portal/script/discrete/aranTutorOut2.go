package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorOut2 struct {
}

func (p AranTutorOut2) Name() string {
	return "aranTutorOut2"
}

func (p AranTutorOut2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {

	processor.TeachSkill(l, c)(20000014, 0, -1, -1)
	processor.TeachSkill(l, c)(20000015, 0, -1, -1)
	processor.TeachSkill(l, c)(20000014, 1, 0, -1)
	processor.TeachSkill(l, c)(20000015, 1, 0, -1)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(914000210, 1)
	return true
}
