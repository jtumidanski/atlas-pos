package properties

import (
	"atlas-pos/model"
	"atlas-pos/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(id uint32) model.Provider[Model] {
	return func(id uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestById(id), makeModel)
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (Model, error) {
	return func(characterId uint32) (Model, error) {
		return ByIdModelProvider(l, span)(characterId)()
	}
}

func ByAccountAndWorldModelListProvider(l logrus.FieldLogger, span opentracing.Span) func(accountId uint32, worldId byte) model.SliceProvider[Model] {
	return func(accountId uint32, worldId byte) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(requestByAccountAndWorld(accountId, worldId), makeModel)
	}
}

func GetForAccountInWorld(l logrus.FieldLogger, span opentracing.Span) func(accountId uint32, worldId byte) ([]Model, error) {
	return func(accountId uint32, worldId byte) ([]Model, error) {
		return ByAccountAndWorldModelListProvider(l, span)(accountId, worldId)()
	}
}

func makeModel(ca requests.DataBody[attributes]) (Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return Model{}, err
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
	return r, nil
}
