package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CatPriestMap struct {
}

func (p CatPriestMap) Name() string {
	return "catPriest_map"
}

func (p CatPriestMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(925000000, 2)
	return true
}
