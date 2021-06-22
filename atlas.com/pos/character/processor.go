package character

import (
	"atlas-pos/job"
	"atlas-pos/kafka/producers"
	"atlas-pos/portal"
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func PlayPortalSound(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {

	}
}

func GetPropertiesById(characterId uint32) (*Properties, error) {
	cs, err := requestAttributesById(characterId)
	if err != nil {
		return nil, err
	}
	ca := makeCharacterAttributes(cs.Data())
	if ca == nil {
		return nil, errors.New("unable to make character attributes")
	}
	return ca, nil
}

func GetAccountCharacters(accountId uint32, worldId byte) ([]*Properties, error) {
	cs, err := requestAccountCharacters(accountId, worldId)
	if err != nil {
		return nil, err
	}

	var cas = make([]*Properties, 0)
	for _, v := range cs.DataList() {
		cas = append(cas, makeCharacterAttributes(&v))
	}
	if len(cas) == 0 {
		return nil, errors.New("unable to make character attributes")
	}
	return cas, nil
}

func ShowInstruction(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) {
	return func(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) {
		err := requests.WorldChannel().CreateInstruction(worldId, channelId, characterId, message, width, height)
		if err != nil {
			l.WithError(err).Errorf("Sending message %s to character %d in world %d channel %d.", message, characterId, worldId, channelId)
		}
	}
}

func HasLevel30Character(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		p, err := GetPropertiesById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character information for character %d.", characterId)
			return false
		}
		cs, err := GetAccountCharacters(p.AccountId(), p.WorldId())
		for _, p = range cs {
			if p.Level() >= 30 {
				return true
			}
		}
		return false
	}
}

func makeCharacterAttributes(ca *attributes.CharacterAttributesData) *Properties {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := NewCharacterAttributeBuilder().
		SetId(uint32(cid)).
		SetAccountId(att.AccountId).
		SetWorldId(att.WorldId).
		SetName(att.Name).
		SetGender(att.Gender).
		SetSkinColor(att.SkinColor).
		SetFace(att.Face).
		SetHair(att.Hair).
		SetLevel(att.Level).
		SetJobId(att.JobId).
		SetStrength(att.Strength).
		SetDexterity(att.Dexterity).
		SetIntelligence(att.Intelligence).
		SetLuck(att.Luck).
		SetHp(att.Hp).
		SetMaxHp(att.MaxHp).
		SetMp(att.Mp).
		SetMaxMp(att.MaxMp).
		SetAp(att.Ap).
		SetSp(att.Sp).
		SetExperience(att.Experience).
		SetFame(att.Fame).
		SetGachaponExperience(att.GachaponExperience).
		SetMapId(att.MapId).
		SetSpawnPoint(att.SpawnPoint).
		SetMeso(att.Meso).
		SetX(att.X).
		SetY(att.Y).
		SetStance(att.Stance).
		Build()
	return &r
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

func WarpToPortal(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
		producers.ChangeMap(l)(worldId, channelId, characterId, mapId, p())
	}
}

func WarpRandom(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
		WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.RandomPortalIdProvider(l)(mapId))
	}
}

func WarpById(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.FixedPortalIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
		WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.ByNamePortalIdProvider(l)(mapId, portalName))
	}
}

func EnableActions(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32) {
	return func(worldId byte, channelId byte, characterId uint32) {
		producers.EnableActions(l)(characterId)
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

type AttributeCriteria func(*Properties) bool

func MeetsCriteria(l logrus.FieldLogger) func(characterId uint32, criteria ...AttributeCriteria) bool {
	return func(characterId uint32, criteria ...AttributeCriteria) bool {
		c, err := GetPropertiesById(characterId)
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

func IsJob(l logrus.FieldLogger) func(characterId uint32, option uint16) bool {
	return func(characterId uint32, option uint16) bool {
		return MeetsCriteria(l)(characterId, IsJobCriteria(option))
	}
}

func IsJobCriteria(option uint16) AttributeCriteria {
	return func(c *Properties) bool {
		return job.IsA(c.JobId(), option)
	}
}

func IsJobNiche(option uint16) AttributeCriteria {
	return func(c *Properties) bool {
		//TODO
		return job.IsA(c.JobId(), option)
	}
}

func RemoveAll(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		//TODO
	}
}

func GetGender(l logrus.FieldLogger) func(characterId uint32) byte {
	return func(characterId uint32) byte {
		c, err := GetPropertiesById(characterId)
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