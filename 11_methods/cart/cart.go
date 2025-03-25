package cart

import "fmt"

type Cart struct {
	Name string
	Ccv  int
}

type money struct {
	Money float32
}

func (c Cart) Display() {
	fmt.Printf("name: %v, ccv: %v\n", c.Name, c.Ccv)
}

func (c *Cart) Change(n string, ccv int) {
	c.Name = n
	c.Ccv = ccv
}

func (c *Cart) TotalPrice() (money, error) {
	if true {
		return money{}, fmt.Errorf("invalid cart")
	}
	return money{100.00}, nil
}
