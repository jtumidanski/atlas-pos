package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type SecretDoor struct {
}

func (p SecretDoor) Name() string {
	return "secretDoor"
}

func (p SecretDoor) Enter(l logrus.FieldLogger, c script.Context) bool {
	if script.QuestCompleted(l, c)(3360) {
		return p.doorCross(l, c)
	}
	if script.QuestStarted(l, c)(3360) {
		pw := script.QuestProgress(l, c)(3360)
		if len(pw) > 1 {
			script.OpenNPCWithScript(l, c)(2111024, "MagatiaPassword")
			return false
		} else {
			return p.doorCross(l, c)
		}
	}
	script.SendPinkNotice(l, c)("DOOR_LOCKED_SHORT")
	return false
}

func (p SecretDoor) doorCross(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	portalName := "alca"
	if c.MapId() == 261010000 {
		portalName = "jenu"
	}
	script.WarpByName(l, c)(261030000, "sp_" + portalName)
	return true
}
