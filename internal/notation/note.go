package notes

type Note struct {
	PitchWithOctave     string
	QuarterNoteDuration int
}

func createTrainingData() []Note {
	return []Note{
		Note{"D4", 1},
		Note{"E4", 1},
		Note{"C4", 2},
		Note{"D4", 1},
		Note{"E4", 1},
		Note{"C4", 2},
		Note{"B4", 1},
		Note{"C4", 1},
		Note{"C4", 1},
		Note{"B4", 1},
		Note{"A4", 2},
	}
}
