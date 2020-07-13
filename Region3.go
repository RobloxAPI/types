package types

type Region3 struct {
	Min, Max Vector3
}

func (r Region3) CFrame() CFrame {
	cf := NewCFrame()
	cf.Position = r.Min.Add(r.Max).DivN(2)
	return cf
}

func (r Region3) Size() Vector3 {
	return r.Max.Sub(r.Min)
}

func (r Region3) ExpandToGrid(res int) Region3 {
	panic("not implemented")
}
