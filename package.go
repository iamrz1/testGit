package main

import (
	"fmt"
	m "maths"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Avg(xs)
	fmt.Println(avg)
	fmt.Println("Nice")
}
