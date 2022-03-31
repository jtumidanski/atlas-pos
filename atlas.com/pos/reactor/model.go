package reactor

type Model struct {
	id    uint32
	state int8
}

func (r Model) State() int8 {
	return r.state
}

func (r Model) Id() uint32 {
	return r.id
}
