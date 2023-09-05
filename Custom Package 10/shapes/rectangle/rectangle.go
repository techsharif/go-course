package rectangle

type Rectangle struct {
	Lenght, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Lenght * r.Width
}
