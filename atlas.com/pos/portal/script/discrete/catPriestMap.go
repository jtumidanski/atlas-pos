package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CatPriestMap struct {
}

func (p CatPriestMap) Name() string {
	return "catPriest_map"
}

func (p CatPriestMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(925000000, 2)
	return true
}
