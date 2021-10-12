package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantAgit struct {
}

func (p AriantAgit) Name() string {
	return "ariant_Agit"
}

func (p AriantAgit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {

	if script.QuestCompleted(l, c)(3928) && script.QuestCompleted(l, c)(3931) && script.QuestCompleted(l, c)(3934) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(260000201, 1)
		return true
	} else {
		script.SendPinkNotice(l, c)("SAND_BANDITS_ONLY")
		return false
	}
}
