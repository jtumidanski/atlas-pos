package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type SecretDoor struct {
}

func (p SecretDoor) Name() string {
	return "secretDoor"
}

func (p SecretDoor) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(3360) {
		return p.doorCross(l, span, c)
	}
	if processor.QuestStarted(l, c)(3360) {
		pw := processor.QuestProgress(l, c)(3360)
		if len(pw) > 1 {
			processor.OpenNPCWithScript(l, c)(2111024, "MagatiaPassword")
			return false
		} else {
			return p.doorCross(l, span, c)
		}
	}
	processor.SendPinkNotice(l, c)("DOOR_LOCKED_SHORT")
	return false
}

func (p SecretDoor) doorCross(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	portalName := "alca"
	if c.MapId() == 261010000 {
		portalName = "jenu"
	}
	processor.WarpByName(l, span, c)(261030000, "sp_" + portalName)
	return true
}
