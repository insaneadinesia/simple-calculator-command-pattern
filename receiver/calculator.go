package receiver

import (
	"fmt"
	"math"
)

type Calculator struct {
	X int
	Y int
}

func (c *Calculator) Sum() {
	total := c.X + c.Y
	fmt.Printf("%d + %d = %d\n", c.X, c.Y, total)
}

func (c *Calculator) Multiply() {
	total := c.X * c.Y
	fmt.Printf("%d x %d = %d\n", c.X, c.Y, total)
}

func (c *Calculator) Prime() {
	number := 1
	total := 1

	for total <= c.X {
		if IsPrimeNumber(number) {
			fmt.Print(number, " ")
			number += 1
			total += 1
			continue
		}

		number += 1
	}

	fmt.Println()
}

func IsPrimeNumber(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}

	return value > 1
}
