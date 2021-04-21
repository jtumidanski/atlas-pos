package script

import (
	"atlas-pos/character"
	"atlas-pos/domain"
	"atlas-pos/kafka/producers"
	"atlas-pos/portal/blocked"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

type Interaction struct {
	l logrus.FieldLogger
	c Context
}

var Processor = func(l logrus.FieldLogger, c Context) *Interaction {
	return &Interaction{l, c}
}

func (p *Interaction) HasLevel30Character() bool {
	cp := character.NewProcessor(p.l)
	c, err := cp.GetCharacterAttributesById(p.c.CharacterId())
	if err != nil {
		p.l.WithError(err).Errorf("Unable to retrieve character information for character %d.", p.c.CharacterId())
		return false
	}
	cs, err := cp.GetAccountCharacters(c.AccountId(), c.WorldId())
	for _, c = range cs {
		if c.Level() >= 30 {
			return true
		}
	}
	return false
}

func (p *Interaction) PlayPortalSound() {
	character.NewProcessor(p.l).PlayPortalSound()
}

func (p *Interaction) WarpById(mapId uint32, portalId uint32) {
	character.NewProcessor(p.l).WarpById(mapId, portalId)
}

func (p *Interaction) WarpByName(mapId uint32, portalName string) {
	character.NewProcessor(p.l).WarpByName(mapId, portalName)
}

func (p *Interaction) ShowInstruction(message string, width int16, height int16) {
	character.NewProcessor(p.l).ShowInstruction(p.c.WorldId(), p.c.ChannelId(), p.c.CharacterId(), message, width, height)
}

func (p *Interaction) OpenNPC(npcId uint32) {
	p.l.Infof("call to unhandled OpenNPC for npc %d from character %d.", npcId, p.c.CharacterId())
	//TODO
}

func (p *Interaction) OpenNPCWithScript(npcId uint32, script string) {
	p.l.Infof("call to unhandled OpenNPC for npc %d from character %d.", npcId, p.c.CharacterId())
	//TODO
}

func (p *Interaction) BlockPortal() {
	p.l.Infof("call to unhandled BlockPortal from character %d.", p.c.CharacterId())
	blocked.GetCache().AddBlockedPortal(p.c.characterId, p.c.portal.ScriptName())
	producers.EnableActions(p.l, context.Background()).Emit(p.c.characterId)
}

func (p *Interaction) QuestStarted(questId uint32) bool {
	p.l.Infof("call to unhandled QuestStarted for quest %d from character %d.", questId, p.c.CharacterId())
	//TODO
	return false
}

func (p *Interaction) ShowInfo(info string) {
	p.l.Infof("call to unhandled ShowInfo for info %s from character %d.", info, p.c.CharacterId())
	//TODO
}

func (p *Interaction) GetSavedLocation(name string) (uint32, uint32) {
	p.l.Infof("call to unhandled GetSavedLocation for location %s.", name)
	//TODO
	return 0, 0
}

func (p *Interaction) SendPinkNotice(token string) {
	p.l.Infof("call to unhandled SendPinkNotice.")
	//TODO
}

func (p *Interaction) GetReactor(mapId uint32, reactorName string) (*domain.ReactorModel, error) {
	p.l.Infof("call to unhandled GetReactor for reactor %s in map %d.", reactorName, mapId)
	return nil, errors.New("not implemented")
}

func (p *Interaction) ContainsAreaInfo(areaId uint16, info string) bool {
	p.l.Infof("call to unhandled ContainsAreaInfo.")
	//TODO
	return false
}

func (p *Interaction) UpdateAreaInfo(areaId uint16, info string) {
	p.l.Infof("call to unhandled UpdateAreaInfo.")
	//TODO
}

func (p *Interaction) ShowIntro(path string) {
	p.l.Infof("call to unhandled ShowIntro.")
	//TODO
}

func (p *Interaction) PlaySound(path string) {
	p.l.Infof("call to unhandled PlaySound.")
	//TODO
}

func (p *Interaction) TeachSkill(skillId uint32, level int8, masterLevel int8, expiration int64) {
	p.l.Infof("call to unhandled TeachSkill.")
	//TODO
}

func (p *Interaction) QuestCompleted(questId uint32) bool {
	p.l.Infof("call to unhandled QuestCompleted for quest %d from character %d.", questId, p.c.CharacterId())
	//TODO
	return false
}

func (p *Interaction) HasItem(itemId uint32) bool {
	p.l.Infof("call to unhandled HasItem.")
	//TODO
	return false
}

func (p *Interaction) Morphed(morphId uint32) bool {
	p.l.Infof("call to unhandled Morphed.")
	//TODO
	return false
}

func (p *Interaction) CanHold(itemId uint32, quantity int16) bool {
	p.l.Infof("call to unhandled CanHold.")
	//TODO
	return false
}

func (p *Interaction) GainItem(itemId uint32, quantity int16) {
	p.l.Infof("call to unhandled GainItem.")
	//TODO
}

func (p *Interaction) EarnTitle(message string) {
	p.l.Infof("call to unhandled EarnTitle.")
	//TODO
}

func (p *Interaction) SetDirectionStatus(b bool) {
	//TODO
}

func (p *Interaction) LockUI() {
	//TODO
}

func (p *Interaction) SpawnNPC(npcId uint32, i2 int, i3 int, i4 int, i5 int, b bool) {
	//TODO
}

func (p *Interaction) SetNPCValue(npdId uint32, value string) {
	//TODO
}

func (p *Interaction) UpdateInfo(key string, value string) {
	//TODO
}

func (p *Interaction) SendDirectionInfo(i int, i2 int) {
	//TODO
}

func (p *Interaction) UnlockUI() {
	//TODO
}

func (p *Interaction) GetParty() (*domain.PartyModel, bool) {
	//TODO
	return nil, false
}

func (p *Interaction) PartyLeader() bool {
	//TODO
	return false
}

func (p *Interaction) StartEvent(eventName string, partyId uint32, mapId uint32, state int) error {
	return nil
}

func (p *Interaction) GetHourOfDay() uint32 {
	return 0
}

func (p *Interaction) MapMonsterCount(mapId uint32) int {
	return 0
}

func (p *Interaction) GetEventProperty(key string) string {
	return ""
}
