package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// basic()
	// literals()
	mutate()
}

func basic() {
	var m map[string]Vertex
	fmt.Println(m)
	m = make(map[string]Vertex)
	m["Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Labs"])
}

func literals() {
	var vt = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	fmt.Println(vt)

	for _, v := range vt {
		fmt.Printf("Lat = %f, Long = %f\n", v.Lat, v.Long)
	}
}

func mutate() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present ?", ok)
}
