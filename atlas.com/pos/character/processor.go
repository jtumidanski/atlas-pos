package character

import (
	"atlas-pos/kafka/producers"
	"atlas-pos/party"
	"atlas-pos/portal"
	"atlas-pos/portal/script"
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func PlayPortalSound(_ logrus.FieldLogger) {
}

func ShowInstruction(l logrus.FieldLogger, c script.Context) func(message string, width int16, height int16) {
	return func(message string, width int16, height int16) {
		err := requests.WorldChannel().CreateInstruction(c.WorldId(), c.ChannelId(), c.CharacterId(), message, width, height)
		if err != nil {
			l.WithError(err).Errorf("Sending message %s to character %d in world %d channel %d.", message, c.CharacterId(), c.WorldId(), c.ChannelId())
		}
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

func HasLevel30Character(l logrus.FieldLogger, c script.Context) bool {
	p, err := GetPropertiesById(c.CharacterId())
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve character information for character %d.", c.CharacterId())
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

func WarpById(l logrus.FieldLogger, c script.Context) func(mapId uint32, portalId uint32) {
	return func(mapId uint32, portalId uint32) {
		producers.ChangeMap(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId, portalId)
	}
}

func WarpByName(l logrus.FieldLogger, c script.Context) func(mapId uint32, portalName string) {
	return func(mapId uint32, portalName string) {
		por, err := portal.GetMapPortalByName(mapId, portalName)
		if err != nil {
			l.WithError(err).Errorf("Unable to lookup portal by name.")
			return
		}
		WarpById(l, c)(mapId, por.Id())
	}
}

func SendPinkNotice(l logrus.FieldLogger, c script.Context) func(token string) {
	return func(token string) {
		l.Infof("call to unhandled SendPinkNotice.")
		//TODO
	}
}
func ShowIntro(l logrus.FieldLogger, c script.Context) func(path string) {
	return func(path string) {
		l.Infof("call to unhandled ShowIntro.")
		//TODO
	}
}

func PlaySound(l logrus.FieldLogger, c script.Context) func(path string) {
	return func(path string) {
		l.Infof("call to unhandled PlaySound.")
		//TODO
	}
}

func TeachSkill(l logrus.FieldLogger, c script.Context) func(skillId uint32, level int8, masterLevel int8, expiration int64) {
	return func(skillId uint32, level int8, masterLevel int8, expiration int64) {
		l.Infof("call to unhandled TeachSkill.")
		//TODO
	}
}
func HasItem(l logrus.FieldLogger, c script.Context) func(itemId uint32) bool {
	return func(itemId uint32) bool {
		l.Infof("call to unhandled HasItem.")
		//TODO
		return false
	}
}

func Morphed(l logrus.FieldLogger, c script.Context) func(morphId uint32) bool {
	return func(morphId uint32) bool {
		l.Infof("call to unhandled Morphed.")
		//TODO
		return false
	}
}

func CanHold(l logrus.FieldLogger, c script.Context) func(itemId uint32, quantity int16) bool {
	return func(itemId uint32, quantity int16) bool {
		l.Infof("call to unhandled CanHold.")
		//TODO
		return false
	}
}

func GainItem(l logrus.FieldLogger, c script.Context) func(itemId uint32, quantity int16) {
	return func(itemId uint32, quantity int16) {
		l.Infof("call to unhandled GainItem.")
		//TODO
	}
}

func EarnTitle(l logrus.FieldLogger, c script.Context) func(message string) {
	return func(message string) {
		l.Infof("call to unhandled EarnTitle.")
		//TODO
	}
}

func LockUI(l logrus.FieldLogger, c script.Context) func() {
	return func() {
		//TODO
	}
}

func UnlockUI(l logrus.FieldLogger, c script.Context) {
	//TODO
}

func GetParty(l logrus.FieldLogger, c script.Context) (*party.Model, bool) {
	//TODO
	return nil, false
}

func PartyLeader(l logrus.FieldLogger, c script.Context) bool {
	//TODO
	return false
}