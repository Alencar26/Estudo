package main

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	Logo string
	Size int
}

func (s *Shoe) setLogo(logo string) {
	s.Logo = logo
}

func (s *Shoe) setSize(size int) {
	s.Size = size
}

func (s *Shoe) getLogo() string {
	return s.Logo
}

func (s *Shoe) getSize() int {
	return s.Size
}
