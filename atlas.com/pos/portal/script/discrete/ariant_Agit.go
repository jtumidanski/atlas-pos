package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantAgit struct {
}

func (p AriantAgit) Name() string {
	return "ariant_Agit"
}

func (p AriantAgit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {

	if processor.QuestCompleted(l, c)(3928) && processor.QuestCompleted(l, c)(3931) && processor.QuestCompleted(l, c)(3934) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(260000201, 1)
		return true
	} else {
		processor.SendPinkNotice(l, c)("SAND_BANDITS_ONLY")
		return false
	}
}
