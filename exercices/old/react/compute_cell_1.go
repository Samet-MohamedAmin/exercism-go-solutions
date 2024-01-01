package react

type computeCell1 struct {
	id        *Cell
	state     *int
	cell      *Cell
	compute   func(int) int
	callbacks *calls
}

func (cc computeCell1) getID() *Cell {
	return cc.id
}

func (cc computeCell1) getState() *int {
	return cc.state
}

func (cc computeCell1) getCallbacks() calls {
	return *cc.callbacks
}

func (cc computeCell1) Value() int {
	return cc.compute((*cc.cell).Value())
}

func (cc computeCell1) AddCallback(callback func(int)) Canceler {
	cc.callbacks.callbacks = append(cc.callbacks.callbacks, &callback)
	return canceller{callback: &callback}
}

type canceller struct {
	callback *func(int)
}

func (c canceller) Cancel() {
	*c.callback = nil
}
