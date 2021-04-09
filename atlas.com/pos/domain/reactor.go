package domain

type ReactorModel struct {
	state byte
}

func (r ReactorModel) State() byte {
	return r.state
}
