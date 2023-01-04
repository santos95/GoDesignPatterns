package main

type ak47 struct {
	gun
}

func newAk47() igun {
	return &ak47{
		gun: gun{
			name:  "Ak47 gun",
			power: 4,
		},
	}
}
