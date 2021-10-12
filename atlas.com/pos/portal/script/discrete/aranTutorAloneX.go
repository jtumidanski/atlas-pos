package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorAloneX struct {
}

func (p AranTutorAloneX) Name() string {
	return "aranTutorAloneX"
}

func (p AranTutorAloneX) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(914000100, 1)
	return true
}
