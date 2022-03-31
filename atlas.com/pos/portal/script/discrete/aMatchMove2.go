package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AMatchMove2 struct {
}

func (p AMatchMove2) Name() string {
	return "aMatchMove2"
}

func (p AMatchMove2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(processor.GetSavedLocation(l, c)("MIRROR"))
	return true
}
