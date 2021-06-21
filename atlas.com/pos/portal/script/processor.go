package script

import (
	"atlas-pos/kafka/producers"
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
