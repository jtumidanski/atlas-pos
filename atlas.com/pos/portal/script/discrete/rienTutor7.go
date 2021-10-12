package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor7 struct {
}

func (p RienTutor7) Name() string {
	return "rienTutor7"
}

func (p RienTutor7) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.IsJob(l, span)(c.CharacterId(), job.Legend) && !script.QuestCompleted(l, c)(21014) {
		script.ShowInfoText(l, c)("The town of Rien is to the right. Take the portal on the right and go into town to meet Lilin.")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(140010100, 2)
	return true
}
