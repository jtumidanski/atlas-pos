package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InXmasParty struct {
}

func (p InXmasParty) Name() string {
	return "in_xmas_party"
}

func (p InXmasParty) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.OpenNPC(l, c)(9209100)
	return false
}
