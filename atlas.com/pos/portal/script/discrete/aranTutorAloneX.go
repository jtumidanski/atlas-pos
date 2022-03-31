package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorAloneX struct {
}

func (p AranTutorAloneX) Name() string {
	return "aranTutorAloneX"
}

func (p AranTutorAloneX) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(914000100, 1)
	return true
}
