package processor

import (
	"atlas-pos/character"
	_map "atlas-pos/map"
	"atlas-pos/monster"
	"atlas-pos/party"
	"atlas-pos/portal"
	"atlas-pos/portal/blocked"
	"atlas-pos/portal/script"
	"atlas-pos/reactor"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func OpenNPC(l logrus.FieldLogger, c script.Context) func(npcId uint32) {
	return func(npcId uint32) {
		l.Infof("call to unhandled OpenNPC for npc %d from character %d.", npcId, c.CharacterId())
		//TODO
	}
}

func OpenNPCWithScript(l logrus.FieldLogger, c script.Context) func(npcId uint32, script string) {
	return func(npcId uint32, script string) {
		l.Infof("call to unhandled OpenNPC for npc %d from character %d.", npcId, c.CharacterId())
		//TODO
	}
}

func BlockPortal(l logrus.FieldLogger, span opentracing.Span, c script.Context) {
	l.Infof("call to unhandled BlockPortal from character %d.", c.CharacterId())
	blocked.GetCache().AddBlockedPortal(c.CharacterId(), c.Portal().ScriptName())
	character.EnableActions(l, span)(c.WorldId(), c.ChannelId(), c.CharacterId())
}

func QuestStarted(l logrus.FieldLogger, c script.Context) func(questId uint32) bool {
	return func(questId uint32) bool {
		l.Infof("call to unhandled QuestStarted for quest %d from character %d.", questId, c.CharacterId())
		//TODO
		return false
	}
}

func QuestActive(l logrus.FieldLogger, c script.Context) func(questId uint32) bool {
	return func(questId uint32) bool {
		l.Infof("call to unhandled QuestActive for quest %d from character %d.", questId, c.CharacterId())
		//TODO
		return false
	}
}

func ShowInfo(l logrus.FieldLogger, c script.Context) func(info string) {
	return func(info string) {
		l.Infof("call to unhandled ShowInfo for info %s from character %d.", info, c.CharacterId())
		//TODO
	}
}

func ShowInfoText(l logrus.FieldLogger, c script.Context) func(message string) {
	return func(info string) {
		l.Infof("call to unhandled ShowInfo for info %s from character %d.", info, c.CharacterId())
		//TODO
	}
}

func ShowEffect(l logrus.FieldLogger, c script.Context) func(path string) {
	return func(path string) {
		l.Infof("call to unhandled ShowInfo for info %s from character %d.", path, c.CharacterId())
		//TODO
	}
}

func SaveLocation(l logrus.FieldLogger, c script.Context) func(name string) {
	return func(name string) {

	}
}

func GetSavedLocation(l logrus.FieldLogger, c script.Context) func(name string) (uint32, uint32) {
	return func(name string) (uint32, uint32) {
		l.Infof("call to unhandled GetSavedLocation for location %s.", name)
		//TODO
		return 0, 0
	}
}

func GetMarketPortal(l logrus.FieldLogger, span opentracing.Span, c script.Context) (uint32, uint32) {
	mapId, _ := GetSavedLocation(l, c)("FREE_MARKET")
	portalId := portal.MarketPortalIdProvider(l, span)(mapId)()
	return mapId, portalId
}

func GetReactor(l logrus.FieldLogger, c script.Context) func(mapId uint32, reactorName string) (*reactor.Model, error) {
	return func(mapId uint32, reactorName string) (*reactor.Model, error) {
		l.Infof("call to unhandled GetReactor for reactor %s in map %d.", reactorName, mapId)
		return nil, errors.New("not implemented")
	}
}

func ContainsAreaInfo(l logrus.FieldLogger, c script.Context) func(areaId uint16, info string) bool {
	return func(areaId uint16, info string) bool {
		l.Infof("call to unhandled ContainsAreaInfo.")
		//TODO
		return false
	}
}

func UpdateAreaInfo(l logrus.FieldLogger, c script.Context) func(areaId uint16, info string) {
	return func(areaId uint16, info string) {
		l.Infof("call to unhandled UpdateAreaInfo.")
		//TODO
	}
}

func QuestCompleted(l logrus.FieldLogger, c script.Context) func(questId uint32) bool {
	return func(questId uint32) bool {
		l.Infof("call to unhandled QuestCompleted for quest %d from character %d.", questId, c.CharacterId())
		//TODO
		return false
	}
}

func SetDirectionStatus(l logrus.FieldLogger, c script.Context) func(b bool) {
	return func(b bool) {
		//TODO
	}
}

func SpawnNPC(l logrus.FieldLogger, c script.Context) func(npcId uint32, i2 int, i3 int, i4 int, i5 int, b bool) {
	return func(npcId uint32, i2 int, i3 int, i4 int, i5 int, b bool) {
		//TODO
	}
}

func SetNPCValue(l logrus.FieldLogger, c script.Context) func(npdId uint32, value string) {
	return func(npdId uint32, value string) {
		//TODO
	}
}

func UpdateInfo(l logrus.FieldLogger, c script.Context) func(key string, value string) {
	return func(key string, value string) {
		//TODO
	}
}

func SendDirectionInfo(l logrus.FieldLogger, c script.Context) func(i int, i2 int) {
	return func(i int, i2 int) {
		//TODO
	}
}

func StartEvent(l logrus.FieldLogger, c script.Context) func(eventName string, partyId uint32, mapId uint32, state int) error {
	return func(eventName string, partyId uint32, mapId uint32, state int) error {
		return nil
	}
}

func GetHourOfDay(l logrus.FieldLogger) uint32 {
	return 0
}

func GetEventProperty(l logrus.FieldLogger, c script.Context) func(key string) string {
	return func(key string) string {
		return ""
	}
}

func WarpById(l logrus.FieldLogger, span opentracing.Span, c script.Context) func(mapId uint32, portalId uint32) {
	return func(mapId uint32, portalId uint32) {
		portal.WarpById(l, span)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId, portalId)
	}
}

func WarpByName(l logrus.FieldLogger, span opentracing.Span, c script.Context) func(mapId uint32, portalName string) {
	return func(mapId uint32, portalName string) {
		portal.WarpByName(l, span)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId, portalName)
	}
}

func WarpRandom(l logrus.FieldLogger, span opentracing.Span, c script.Context) func(mapId uint32) {
	return func(mapId uint32) {
		portal.WarpRandom(l, span)(c.WorldId(), c.ChannelId(), c.CharacterId(), mapId)
	}
}

func SendPinkNotice(l logrus.FieldLogger, c script.Context) func(token string) {
	return func(token string) {
		character.SendPinkNotice(l)(c.WorldId(), c.ChannelId(), c.CharacterId(), token)
	}
}

func SendLightBlueNotice(l logrus.FieldLogger, c script.Context) func(token string) {
	return func(token string) {

	}
}

func EnableActions(l logrus.FieldLogger, span opentracing.Span, c script.Context) {
	character.EnableActions(l, span)(c.WorldId(), c.ChannelId(), c.CharacterId())
}

func GainExperience(l logrus.FieldLogger, c script.Context) func(amount int32) {
	return func(amount int32) {
		character.GainExperience(l)(c.CharacterId(), amount)
	}
}

func ShowInstruction(l logrus.FieldLogger, span opentracing.Span, c script.Context) func(message string, width int16, height int16) {
	return func(message string, width int16, height int16) {
		character.ShowInstruction(l, span)(c.WorldId(), c.ChannelId(), c.CharacterId(), message, width, height)
	}
}

func HasLevel30Character(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	return character.HasLevel30Character(l, span)(c.CharacterId())
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

func GetParty(l logrus.FieldLogger, span opentracing.Span, c script.Context) (party.Model, bool) {
	p, err := party.GetByMemberId(l, span)(c.CharacterId())
	if err != nil {
		l.WithError(err).Errorf("Unable to retrieve party for character %d.", c.CharacterId())
		return party.Model{}, false
	}
	return p, false
}

func PartyLeader(l logrus.FieldLogger, c script.Context) bool {
	//TODO
	return false
}

func SetNPCCooldown(l logrus.FieldLogger, c script.Context) func(time int64) {
	return func(time int64) {

	}
}

func NPCCooldown(l logrus.FieldLogger, c script.Context) int64 {
	return 0
}

func DojoPoints(l logrus.FieldLogger, c script.Context) uint32 {
	return DojoPointsOtherCharacter(l, c)(c.CharacterId())
}

func DojoPointsOtherCharacter(l logrus.FieldLogger, c script.Context) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		return 0
	}
}

func AwardDojoPoints(l logrus.FieldLogger, c script.Context) func(amount uint32) {
	return func(amount uint32) {
		AwardDojoPointsOtherCharacter(l, c)(c.CharacterId(), amount)
	}
}

func AwardDojoPointsOtherCharacter(l logrus.FieldLogger, c script.Context) func(characterId uint32, amount uint32) {
	return func(characterId uint32, amount uint32) {

	}
}

func PlayPortalSound(l logrus.FieldLogger, c script.Context) {
	character.PlayPortalSound(l)(c.CharacterId())
}

func MonsterById(l logrus.FieldLogger, c script.Context) func(monsterId uint32) *monster.Model {
	return func(monsterId uint32) *monster.Model {
		return _map.MonsterById(l)(c.WorldId(), c.ChannelId(), c.MapId(), monsterId)
	}
}

func ReactorByName(l logrus.FieldLogger, span opentracing.Span, c script.Context) func(reactorName string) (reactor.Model, error) {
	return func(reactorName string) (reactor.Model, error) {
		return reactor.ByName(l, span)(c.WorldId(), c.ChannelId(), c.MapId(), reactorName)
	}
}

func MonsterCount(l logrus.FieldLogger, c script.Context) int {
	return _map.MonstersCount(l)(c.WorldId(), c.ChannelId(), c.MapId())
}

func CharactersInMap(l logrus.FieldLogger, c script.Context) []uint32 {
	return _map.CharactersInMap(l)(c.WorldId(), c.ChannelId(), c.MapId())
}

func QuestProgressInt(l logrus.FieldLogger, c script.Context) func(questId uint32, infoNumber int) int {
	return func(questId uint32, infoNumber int) int {
		return character.QuestProgressInt(l)(c.CharacterId(), questId, infoNumber)
	}
}

func SetQuestProgress(l logrus.FieldLogger, c script.Context) func(questId uint32, infoNumber int, progress uint32) {
	return func(questId uint32, infoNumber int, progress uint32) {
		character.SetQuestProgress(l)(c.CharacterId(), questId, infoNumber, progress)
	}
}

func RunMapScript(l logrus.FieldLogger, c script.Context) {

}

func MapEffect(l logrus.FieldLogger, c script.Context) func(path string) {
	return func(path string) {

	}
}

func UseItem(l logrus.FieldLogger, c script.Context) func(itemId uint32) {
	return func(itemId uint32) {
		character.UseItem(l)(c.CharacterId(), itemId)
	}
}

func SpawnGuide(l logrus.FieldLogger, c script.Context) {
}

func TalkGuide(l logrus.FieldLogger, c script.Context) func(message string) {
	return func(message string) {

	}
}

func RemoveGuide(l logrus.FieldLogger, c script.Context) {

}

func QuestProgress(l logrus.FieldLogger, c script.Context) func(questId uint32) string {
	return func(questId uint32) string {
		return character.QuestProgress(l)(c.CharacterId(), questId)
	}
}

func HasItems(l logrus.FieldLogger, c script.Context) func(itemId uint32, amount int16) bool {
	return func(itemId uint32, amount int16) bool {
		l.Infof("call to unhandled HasItem.")
		//TODO
		return false
	}
}

func CancelItem(l logrus.FieldLogger, c script.Context) func(itemId uint32) {
	return func(itemId uint32) {

	}
}

func GuideHint(l logrus.FieldLogger, c script.Context) func(hint uint32) {
	return func(hint uint32) {
		character.GuideHint(l)(c.CharacterId(), hint)
	}
}
