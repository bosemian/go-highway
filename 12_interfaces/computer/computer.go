package computer

type Computer struct {
	Storage float64
}

func (c *Computer) AddStorage() {
	c.Storage += 100.00
}
