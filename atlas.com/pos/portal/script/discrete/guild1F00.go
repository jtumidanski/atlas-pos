package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Guild1F00 struct {
}

func (p Guild1F00) Name() string {
	return "guild1F00"
}

func (p Guild1F00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//int[] backPortals = [6, 8, 9, 11]
	//int idx = pi.getEventInstance().gridCheck(pi.getPlayer())
	//
	processor.PlayPortalSound(l, c)
	//pi.warp(990000600, backPortals[idx])
	return true
}
