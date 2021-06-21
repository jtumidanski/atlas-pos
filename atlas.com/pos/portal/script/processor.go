package script

import (
	"atlas-pos/character"
	"atlas-pos/kafka/producers"
	_map "atlas-pos/map"
	"atlas-pos/monster"
	"atlas-pos/party"
	"atlas-pos/portal/blocked"
	"atlas-pos/reactor"
	"errors"
	"github.com/sirupsen/logrus"
)

func OpenNPC(l logrus.FieldLogger, c Context) func(npcId uint32) {
	return func(npcId uint32) {
		l.Infof("call to unhandled OpenNPC for npc %d from character %d.", npcId, c.CharacterId())
		//TODO
	}
}

func OpenNPCWithScript(l logrus.FieldLogger, c Context) func(npcId uint32, script string) {
	return func(npcId uint32, script string) {
		l.Infof("call to unhandled OpenNPC for npc %d from character %d.", npcId, c.CharacterId())
		//TODO
	}
}

func BlockPortal(l logrus.FieldLogger, c Context) {
	l.Infof("call to unhandled BlockPortal from character %d.", c.CharacterId())
	blocked.GetCache().AddBlockedPortal(c.CharacterId(), c.portal.ScriptName())
	producers.EnableActions(l)(c.CharacterId())
}

func QuestStarted(l logrus.FieldLogger, c Context) func(questId uint32) bool {
	return func(questId uint32) bool {
		l.Infof("call to unhandled QuestStarted for quest %d from character %d.", questId, c.CharacterId())
		//TODO
		return false
	}
}

func ShowInfo(l logrus.FieldLogger, c Context) func(info string) {
	return func(info string) {
		l.Infof("call to unhandled ShowInfo for info %s from character %d.", info, c.CharacterId())
		//TODO
	}
}

func GetSavedLocation(l logrus.FieldLogger, c Context) func(name string) (uint32, uint32) {
	return func(name string) (uint32, uint32) {
		l.Infof("call to unhandled GetSavedLocation for location %s.", name)
		//TODO
		return 0, 0
	}
}

func GetReactor(l logrus.FieldLogger, c Context) func(mapId uint32, reactorName string) (*reactor.Model, error) {
	return func(mapId uint32, reactorName string) (*reactor.Model, error) {
		l.Infof("call to unhandled GetReactor for reactor %s in map %d.", reactorName, mapId)
		return nil, errors.New("not implemented")
	}
}

func ContainsAreaInfo(l logrus.FieldLogger, c Context) func(areaId uint16, info string) bool {
	return func(areaId uint16, info string) bool {
		l.Infof("call to unhandled ContainsAreaInfo.")
		//TODO
		return false
	}
}

func UpdateAreaInfo(l logrus.FieldLogger, c Context) func(areaId uint16, info string) {
	return func(areaId uint16, info string) {
		l.Infof("call to unhandled UpdateAreaInfo.")
		//TODO
	}
}

func QuestCompleted(l logrus.FieldLogger, c Context) func(questId uint32) bool {
	return func(questId uint32) bool {
		l.Infof("call to unhandled QuestCompleted for quest %d from character %d.", questId, c.CharacterId())
		//TODO
		return false
	}
}

func SetDirectionStatus(l logrus.FieldLogger, c Context) func(b bool) {
	return func(b bool) {
		//TODO
	}
}

func SpawnNPC(l logrus.FieldLogger, c Context) func(npcId uint32, i2 int, i3 int, i4 int, i5 int, b bool) {
	return func(npcId uint32, i2 int, i3 int, i4 int, i5 int, b bool) {
		//TODO
	}
}

func SetNPCValue(l logrus.FieldLogger, c Context) func(npdId uint32, value string) {
	return func(npdId uint32, value string) {
		//TODO
	}
}

func UpdateInfo(l logrus.FieldLogger, c Context) func(key string, value string) {
	return func(key string, value string) {
		//TODO
	}
}

func SendDirectionInfo(l logrus.FieldLogger, c Context) func(i int, i2 int) {
	return func(i int, i2 int) {
		//TODO
	}
}

func StartEvent(l logrus.FieldLogger, c Context) func(eventName string, partyId uint32, mapId uint32, state int) error {
	return func(eventName string, partyId uint32, mapId uint32, state int) error {
		return nil
	}
}

func GetHourOfDay(l logrus.FieldLogger) uint32 {
	return 0
}

func GetEventProperty(l logrus.FieldLogger, c Context) func(key string) string {
	return func(key string) string {
		return ""
	}
}

func WarpById(l logrus.FieldLogger, c Context) func(mapId uint32, portalId uint32) {
	return func(mapId uint32, portalId uint32) {
		character.WarpById(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId, portalId)
	}
}

func WarpByName(l logrus.FieldLogger, c Context) func(mapId uint32, portalName string) {
	return func(mapId uint32, portalName string) {
		character.WarpByName(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId, portalName)
	}
}

func WarpRandom(l logrus.FieldLogger, c Context) func(mapId uint32) {
	return func(mapId uint32) {
		character.WarpRandom(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId)
	}
}

func SendPinkNotice(l logrus.FieldLogger, c Context) func(token string) {
	return func(token string) {
		character.SendPinkNotice(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), token)
	}
}

func EnableActions(l logrus.FieldLogger, c Context) {
	character.EnableActions(l)(c.WorldId(), c.ChannelId(), c.CharacterId())
}

func GainExperience(l logrus.FieldLogger, c Context) func(amount int32) {
	return func(amount int32) {
		character.GainExperience(l)(c.CharacterId(), amount)
	}
}

func ShowInstruction(l logrus.FieldLogger, c Context) func(message string, width int16, height int16) {
	return func(message string, width int16, height int16) {
		character.ShowInstruction(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), message, width, height)
	}
}

func HasLevel30Character(l logrus.FieldLogger, c Context) bool {
	return character.HasLevel30Character(l)(c.CharacterId())
}

func ShowIntro(l logrus.FieldLogger, c Context) func(path string) {
	return func(path string) {
		l.Infof("call to unhandled ShowIntro.")
		//TODO
	}
}

func PlaySound(l logrus.FieldLogger, c Context) func(path string) {
	return func(path string) {
		l.Infof("call to unhandled PlaySound.")
		//TODO
	}
}

func TeachSkill(l logrus.FieldLogger, c Context) func(skillId uint32, level int8, masterLevel int8, expiration int64) {
	return func(skillId uint32, level int8, masterLevel int8, expiration int64) {
		l.Infof("call to unhandled TeachSkill.")
		//TODO
	}
}
func HasItem(l logrus.FieldLogger, c Context) func(itemId uint32) bool {
	return func(itemId uint32) bool {
		l.Infof("call to unhandled HasItem.")
		//TODO
		return false
	}
}

func Morphed(l logrus.FieldLogger, c Context) func(morphId uint32) bool {
	return func(morphId uint32) bool {
		l.Infof("call to unhandled Morphed.")
		//TODO
		return false
	}
}

func CanHold(l logrus.FieldLogger, c Context) func(itemId uint32, quantity int16) bool {
	return func(itemId uint32, quantity int16) bool {
		l.Infof("call to unhandled CanHold.")
		//TODO
		return false
	}
}

func GainItem(l logrus.FieldLogger, c Context) func(itemId uint32, quantity int16) {
	return func(itemId uint32, quantity int16) {
		l.Infof("call to unhandled GainItem.")
		//TODO
	}
}

func EarnTitle(l logrus.FieldLogger, c Context) func(message string) {
	return func(message string) {
		l.Infof("call to unhandled EarnTitle.")
		//TODO
	}
}

func LockUI(l logrus.FieldLogger, c Context) func() {
	return func() {
		//TODO
	}
}

func UnlockUI(l logrus.FieldLogger, c Context) {
	//TODO
}

func GetParty(l logrus.FieldLogger, c Context) (*party.Model, bool) {
	//TODO
	return nil, false
}

func PartyLeader(l logrus.FieldLogger, c Context) bool {
	//TODO
	return false
}

func SetNPCCooldown(l logrus.FieldLogger, c Context) func(time int64) {
	return func(time int64) {

	}
}

func NPCCooldown(l logrus.FieldLogger, c Context) int64 {
	return 0
}

func DojoPoints(l logrus.FieldLogger, c Context) uint32 {
	return DojoPointsOtherCharacter(l, c)(c.CharacterId())
}

func DojoPointsOtherCharacter(l logrus.FieldLogger, c Context) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		return 0
	}
}

func AwardDojoPoints(l logrus.FieldLogger, c Context) func(amount uint32) {
	return func(amount uint32) {
		AwardDojoPointsOtherCharacter(l, c)(c.CharacterId(), amount)
	}
}

func AwardDojoPointsOtherCharacter(l logrus.FieldLogger, c Context) func(characterId uint32, amount uint32) {
	return func(characterId uint32, amount uint32) {

	}
}

func PlayPortalSound(l logrus.FieldLogger, c Context) {
	character.PlayPortalSound(l)(c.CharacterId())
}

func MonsterById(l logrus.FieldLogger, c Context) func(monsterId uint32) *monster.Model {
	return func(monsterId uint32) *monster.Model {
		return _map.MonsterById(l)(c.WorldId(), c.ChannelId(), c.MapId(), monsterId)
	}
}

func ReactorByName(l logrus.FieldLogger, c Context) func(reactorName string) *reactor.Model {
	return func(reactorName string) *reactor.Model {
		return reactor.ByName(l)(c.WorldId(), c.ChannelId(), c.MapId(), reactorName)
	}
}

func MonsterCount(l logrus.FieldLogger, c Context) int {
	return _map.MonsterCount(l)(c.WorldId(), c.ChannelId(), c.MapId())
}

func CharactersInMap(l logrus.FieldLogger, c Context) []uint32 {
	return _map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), c.MapId())
}
