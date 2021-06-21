package reactor

type Model struct {
	state byte
}

func (r Model) State() byte {
	return r.state
}
