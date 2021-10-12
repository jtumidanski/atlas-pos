package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorOut2 struct {
}

func (p AranTutorOut2) Name() string {
	return "aranTutorOut2"
}

func (p AranTutorOut2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {

	script.TeachSkill(l, c)(20000014, 0, -1, -1)
	script.TeachSkill(l, c)(20000015, 0, -1, -1)
	script.TeachSkill(l, c)(20000014, 1, 0, -1)
	script.TeachSkill(l, c)(20000015, 1, 0, -1)
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(914000210, 1)
	return true
}
