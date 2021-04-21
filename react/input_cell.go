package react

type inputCell struct {
	id    *Cell
	value *int
	r     reactor
}

func (ic inputCell) getID() *Cell {
	return ic.id
}

func (ic inputCell) Value() int {
	return *ic.value
}

func (ic inputCell) SetValue(value int) {
	oldValue := ic.Value()
	if value != oldValue {
		*ic.value = value
		ic.r.valueChanged(ic.id)
	}
}
