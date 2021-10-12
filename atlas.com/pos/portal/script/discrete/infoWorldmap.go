package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type InfoWorldMap struct {
}

func (p InfoWorldMap) Name() string {
	return "infoWorldmap"
}

func (p InfoWorldMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.ShowInfo(l, c)("UI/tutorial.img/26")
	script.BlockPortal(l, span, c)
	return true
}
