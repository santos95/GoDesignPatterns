package main

type iShoe interface {
	setSize(size int)
	setLogo(logo string)
	getSize() int
	getLogo() string
}

type shoe struct {
	size int
	logo string
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getSize() int {
	return s.size
}

func (s *shoe) getLogo() string {
	return s.logo
}
