package receiver

import "fmt"

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
