package main

import notes "github.com/guitarkeegan/markov-go/internal/notation"

func createTrainingData() []notes.Note {

	return []notes.Note{
		{
			PitchWithOctave:     "G4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "C5",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "B4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "A4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "G4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "A4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "B4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "D5",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "B4",
			QuarterNoteDuration: 1,
		},
		{
			PitchWithOctave:     "C5",
			QuarterNoteDuration: 2,
		},
	}
}
