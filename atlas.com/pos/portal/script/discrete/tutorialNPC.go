package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type TutorialNPC struct {
}

func (p TutorialNPC) Name() string {
	return "tutorialNPC"
}

func (p TutorialNPC) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.AboveLevel(l)(c.CharacterId(), 10) || !character.IsJob(l)(c.CharacterId(), job.Beginner) {
		return false
	}

	var npcId uint32
	if c.MapId() == 120000101 {
		npcId = 1090000
	} else if c.MapId() == 102000003 {
		npcId = 1022000
	} else if c.MapId() == 103000003 {
		npcId = 1052001
	} else if c.MapId() == 100000201 {
		npcId = 1012100
	} else if c.MapId() == 101000003 {
		npcId = 1032001
	}

	if npcId == 0 {
		return false
	}

	script.OpenNPC(l, c)(npcId)
	return true
}
