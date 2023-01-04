package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{shoe{
		logo: "nike",
		size: 8,
	}}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 28,
		},
	}
}
