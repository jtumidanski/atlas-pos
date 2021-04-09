package domain

type PartyModel struct {
	id uint32
}

func (m PartyModel) Id() uint32 {
	return m.id
}
