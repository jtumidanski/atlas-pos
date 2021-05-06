package character

import (
	"atlas-pos/domain"
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Processor struct {
	l logrus.FieldLogger
}

func NewProcessor(l logrus.FieldLogger) *Processor {
	return &Processor{l}
}

func (p *Processor) PlayPortalSound() {
}

func (p *Processor) ShowInstruction(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) {
	err := requests.WorldChannel().CreateInstruction(worldId, channelId, characterId, message, width, height)
	if err != nil {
		p.l.WithError(err).Errorf("Sending message %s to character %d in world %d channel %d.", message, characterId, worldId, channelId)
	}
}

func (p *Processor) GetCharacterAttributesById(characterId uint32) (*domain.CharacterAttributes, error) {
	cs, err := requests.Character().GetCharacterAttributesById(characterId)
	if err != nil {
		return nil, err
	}
	ca := makeCharacterAttributes(cs.Data())
	if ca == nil {
		return nil, errors.New("unable to make character attributes")
	}
	return ca, nil
}

func (p *Processor) GetAccountCharacters(accountId uint32, worldId byte) ([]*domain.CharacterAttributes, error) {
	cs, err := requests.Character().GetAccountCharacters(accountId, worldId)
	if err != nil {
		return nil, err
	}

	var cas = make([]*domain.CharacterAttributes, 0)
	for _, v := range cs.DataList() {
		cas = append(cas, makeCharacterAttributes(&v))
	}
	if len(cas) == 0 {
		return nil, errors.New("unable to make character attributes")
	}
	return cas, nil
}

func makeCharacterAttributes(ca *attributes.CharacterAttributesData) *domain.CharacterAttributes {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := domain.NewCharacterAttributeBuilder().
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
