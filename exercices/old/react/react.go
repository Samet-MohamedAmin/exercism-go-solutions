package react

type reactor struct {
	cellsMap map[*Cell][]*Cell
}


func (r reactor) valueChanged(id *Cell) {
	for _, c := range r.cellsMap[id] {
		switch c2 := (*c).(type) {
		case ComputeCell:
			if old, new := c2.getState(), c2.Value(); new != *old {
				*old = new
				for _, callback := range c2.getCallbacks().callbacks {
					if (*callback) != nil {
						(*callback)(new)
					}
				}
				r.valueChanged(c)
			}
		}
	}
}

// New creates new reactor
func New() Reactor {
	return reactor{map[*Cell][]*Cell{}}
}

func (r reactor) CreateInput(value int) InputCell {
	ic := inputCell{value: &value, r: r}
	k := Cell(ic)
	ic.id = &k
	return ic
}

func (r reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cl := calls{}
	state := f(c.Value())
	cc := computeCell1{cell: c.getID(), compute: f, callbacks: &cl, state: &state}
	k := Cell(cc)
	cc.id = &k

	r.cellsMap[c.getID()] = append(r.cellsMap[c.getID()], &k)
	return cc
}

func (r reactor) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	cl := calls{}
	state := f(c1.Value(), c2.Value())
	cc := computeCell2{cell1: c1.getID(), cell2: c2.getID(), compute: f, callbacks: &cl, state: &state}
	k := Cell(cc)
	cc.id = &k

	r.cellsMap[c1.getID()] = append(r.cellsMap[c1.getID()], &k)
	r.cellsMap[c2.getID()] = append(r.cellsMap[c2.getID()], &k)
	return cc
}
