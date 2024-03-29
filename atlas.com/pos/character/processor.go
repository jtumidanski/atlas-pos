package character

import (
	"atlas-pos/character/properties"
	"atlas-pos/instruction"
	"atlas-pos/job"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func PlayPortalSound(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {

	}
}

func ShowInstruction(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) {
	return func(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) {
		_, _, err := instruction.Create(worldId, channelId, characterId, message, width, height)(l, span)
		if err != nil {
			l.WithError(err).Errorf("Sending message %s to character %d in world %d channel %d.", message, characterId, worldId, channelId)
		}
	}
}

func HasLevel30Character(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		p, err := properties.GetById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character information for character %d.", characterId)
			return false
		}
		cs, err := properties.GetForAccountInWorld(l, span)(p.AccountId(), p.WorldId())
		for _, p = range cs {
			if p.Level() >= 30 {
				return true
			}
		}
		return false
	}
}

func SendPinkNotice(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, token string) {
	return func(worldId byte, channelId byte, characterId uint32, token string) {
		//TODO
	}
}

func GainExperience(l logrus.FieldLogger) func(characterId uint32, amount int32) {
	return func(characterId uint32, amount int32) {
		//TODO
	}
}

func ForceCompleteQuestViaNPC(l logrus.FieldLogger) func(characterId uint32, questId uint32, npcId uint32) {
	return func(characterId uint32, questId uint32, npcId uint32) {
		//TODO
	}
}

func ForceCompleteQuest(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {
		//TODO
	}
}

func ForceStartQuest(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {
		//TODO
	}
}

func EnableActions(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32) {
	return func(worldId byte, channelId byte, characterId uint32) {
		emitEnableActions(l, span)(characterId)
	}
}

func QuestProgressInt(l logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber int) int {
	return func(characterId uint32, questId uint32, infoNumber int) int {
		//TODO
		return 0
	}
}

func SetQuestProgress(l logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber int, progress uint32) {
	return func(characterId uint32, questId uint32, infoNumber int, progress uint32) {
		//TODO
	}
}

type PropertiesCriteria func(properties.Model) bool

func MeetsCriteria(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, criteria ...PropertiesCriteria) bool {
	return func(characterId uint32, criteria ...PropertiesCriteria) bool {
		c, err := properties.GetById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d for criteria check.", characterId)
			return false
		}
		for _, check := range criteria {
			if ok := check(c); !ok {
				return false
			}
		}
		return true
	}
}

func IsJob(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, option uint16) bool {
	return func(characterId uint32, option uint16) bool {
		return MeetsCriteria(l, span)(characterId, IsJobCriteria(option))
	}
}

func IsJobCriteria(option uint16) PropertiesCriteria {
	return func(c properties.Model) bool {
		return job.IsA(c.JobId(), option)
	}
}

func IsJobNiche(option uint16) PropertiesCriteria {
	return func(c properties.Model) bool {
		//TODO
		return job.IsA(c.JobId(), option)
	}
}

func RemoveAll(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		//TODO
	}
}

func GetGender(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) byte {
	return func(characterId uint32) byte {
		c, err := properties.GetById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Gender()
	}
}

func UseItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		//TODO
	}
}

func ItemQuantity(l logrus.FieldLogger) func(characterId uint32, itemId uint32) uint32 {
	return func(characterId uint32, itemId uint32) uint32 {
		//TODO
		return 0
	}
}

func QuestProgress(l logrus.FieldLogger) func(characterId uint32, questId uint32) string {
	return func(characterId uint32, questId uint32) string {
		//TODO
		return ""
	}
}

func AboveLevel(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, level byte) bool {
	return func(characterId uint32, level byte) bool {
		return MeetsCriteria(l, span)(characterId, AboveLevelCriteria(level))
	}
}

func AboveLevelCriteria(level byte) PropertiesCriteria {
	return func(c properties.Model) bool {
		return c.Level() > level
	}
}

func GuideHint(l logrus.FieldLogger) func(characterId uint32, hint uint32) {
	return func(characterId uint32, hint uint32) {
		//TODO
	}
}
