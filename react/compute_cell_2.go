package react

type computeCell2 struct {
	id        *Cell
	state     *int
	cell1     *Cell
	cell2     *Cell
	compute   func(int, int) int
	callbacks *calls
}

func (cc computeCell2) getID() *Cell {
	return cc.id
}

func (cc computeCell2) getState() *int {
	return cc.state
}

func (cc computeCell2) getCallbacks() calls {
	return *cc.callbacks
}

func (cc computeCell2) Value() int {
	v1 := (*cc.cell1).Value()
	v2 := (*cc.cell2).Value()
	return cc.compute(v1, v2)
}

func (cc computeCell2) AddCallback(callback func(int)) Canceler {
	cc.callbacks.callbacks = append(cc.callbacks.callbacks, &callback)
	return canceller{callback: &callback}
}
