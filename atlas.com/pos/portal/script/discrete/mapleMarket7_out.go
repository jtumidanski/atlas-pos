package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MapleMarket7Out struct {
}

func (p MapleMarket7Out) Name() string {
	return "mapleMarket7_out"
}

func (p MapleMarket7Out) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpToSavedLocation(l, span, c)("EVENT")
	return true
}
