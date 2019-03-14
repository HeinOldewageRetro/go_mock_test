package tool

type Struct struct {
	I int
}

func (s *Struct) Int() int {
	return s.I
}
