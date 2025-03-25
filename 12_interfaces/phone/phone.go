package phone

type CellPhone struct {
	Brand   string
	Storage float64
}

func (p *CellPhone) AddStorage() {
	p.Storage += 100.00
}
