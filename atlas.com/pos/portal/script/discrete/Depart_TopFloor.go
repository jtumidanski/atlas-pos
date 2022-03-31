package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartTopFloor struct {
}

func (p DepartTopFloor) Name() string {
	return "Depart_TopFloor"
}

func (p DepartTopFloor) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.OpenNPC(l, c)(1052125)
	return true
}
