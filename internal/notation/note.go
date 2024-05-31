package notes

type Note struct {
	PitchWithOctave     string
	QuarterNoteDuration int
}

func createTrainingData() []Note {
	return []Note{
		{"D4", 1},
		{"E4", 1},
		{"C4", 2},
		{"D4", 1},
		{"E4", 1},
		{"C4", 2},
		{"B4", 1},
		{"C4", 1},
		{"C4", 1},
		{"B4", 1},
		{"A4", 2},
	}
}
