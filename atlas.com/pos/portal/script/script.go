package script

import (
	"atlas-pos/domain"
	"log"
)

type Context struct {
	worldId     byte
	channelId   byte
	characterId uint32
	mapId       uint32
	portal      *domain.PortalModel
}

func (c Context) MapId() uint32 {
	return c.mapId
}

func (c Context) WorldId() byte {
	return c.worldId
}

func (c Context) ChannelId() byte {
	return c.channelId
}

func (c Context) CharacterId() uint32 {
	return c.characterId
}

func (c Context) Portal() *domain.PortalModel {
	return c.portal
}

func NewContext(worldId byte, channelId byte, characterId uint32, mapId uint32, portal *domain.PortalModel) Context {
	return Context{
		worldId:     worldId,
		channelId:   channelId,
		characterId: characterId,
		mapId:       mapId,
		portal:      portal,
	}
}

type Script interface {
	Name() string
	AsPortalScript(l *log.Logger, c Context) (*PortalScriptImpl, bool)
}

type ScriptImpl struct {
	si interface{}
}

func NewScript(ps PortalScript) ScriptImpl {
	return ScriptImpl{ps}
}

func (s ScriptImpl) Name() string {
	if val, ok := s.si.(PortalScript); ok {
		return val.Name()
	}
	return ""
}

func (s ScriptImpl) AsPortalScript(l *log.Logger, c Context) (*PortalScriptImpl, bool) {
	if val, ok := s.si.(PortalScript); ok {
		return &PortalScriptImpl{
			l:     l,
			c:     c,
			name:  val.Name,
			enter: val.Enter,
		}, true
	}
	return nil, false
}

type PortalScriptImpl struct {
	l     *log.Logger
	c     Context
	name  func() string
	enter func(l *log.Logger, c Context) bool
}

func (p PortalScriptImpl) Name() string {
	return p.name()
}

func (p PortalScriptImpl) Enter() bool {
	return p.enter(p.l, p.c)
}

type PortalScript interface {
	Name() string
	Enter(l *log.Logger, c Context) bool
}

type PortalScriptName func() string

type PortalScriptEnter func(l *log.Logger, c Context) bool
