package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	area := rect.Area()
	fmt.Println("Area:", area)

	counter := Counter{Count: 0}
	counter.Increment()
	fmt.Println("Count:", counter.Count)
}

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}
