package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type SpaceReturn struct {
}

func (p SpaceReturn) Name() string {
	return "space_return"
}

func (p SpaceReturn) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpToSavedLocation(l, span, c)("EVENT")
	return true
}
