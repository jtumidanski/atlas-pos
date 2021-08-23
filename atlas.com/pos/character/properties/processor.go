package properties

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetById(l logrus.FieldLogger) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		cs, err := requestAttributesById(l)(characterId)
		if err != nil {
			return nil, err
		}
		ca := makeModel(cs.Data())
		if ca == nil {
			return nil, errors.New("unable to make character attributes")
		}
		return ca, nil
	}
}

func GetForAccountInWorld(l logrus.FieldLogger) func(accountId uint32, worldId byte) ([]*Model, error) {
	return func(accountId uint32, worldId byte) ([]*Model, error) {
		cs, err := requestAccountCharacters(l)(accountId, worldId)
		if err != nil {
			return nil, err
		}

		var cas = make([]*Model, 0)
		for _, v := range cs.DataList() {
			cas = append(cas, makeModel(&v))
		}
		if len(cas) == 0 {
			return nil, errors.New("unable to make character attributes")
		}
		return cas, nil
	}
}

func makeModel(ca *CharacterAttributesData) *Model {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := NewBuilder().
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