package location

type Model struct {
	locationType string
	mapId        uint32
	portalId     uint32
}

func (m Model) MapId() uint32 {
	return m.mapId
}

func (m Model) PortalId() uint32 {
	return m.portalId
}
