package main

type IShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	Logo string
	Size int
}

func (s *Shirt) setLogo(logo string) {
	s.Logo = logo
}

func (s *Shirt) setSize(size int) {
	s.Size = size
}

func (s *Shirt) getLogo() string {
	return s.Logo
}

func (s *Shirt) getSize() int {
	return s.Size
}
