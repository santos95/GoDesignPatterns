package main

type musket struct {
	gun
}

func newMusket() igun {
	return &musket{
		gun: gun{
			name:  "musket gun",
			power: 2,
		},
	}

}
