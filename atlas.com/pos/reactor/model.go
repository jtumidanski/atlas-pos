package reactor

type Model struct {
	id uint32
	state byte
}

func (r Model) State() byte {
	return r.state
}

func (r Model) Id() uint32 {
	return r.id
}
