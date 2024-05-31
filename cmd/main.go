package main

import (
	"fmt"
	"log"

	model "github.com/guitarkeegan/markov-go/internal/markov"
	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

func main() {

	// create training data
	trainingData := createTrainingData()
	fmt.Println(trainingData)
	// define a series of states
	states := []notes.Note{
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "F4", QuarterNoteDuration: 1},
		{PitchWithOctave: "G4", QuarterNoteDuration: 1},
		{PitchWithOctave: "A4", QuarterNoteDuration: 1},
		{PitchWithOctave: "B4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 2},
		{PitchWithOctave: "D4", QuarterNoteDuration: 2},
		{PitchWithOctave: "E4", QuarterNoteDuration: 2},
		{PitchWithOctave: "F4", QuarterNoteDuration: 2},
		{PitchWithOctave: "G4", QuarterNoteDuration: 2},
		{PitchWithOctave: "A4", QuarterNoteDuration: 2},
		{PitchWithOctave: "B4", QuarterNoteDuration: 2},
	}
	// initialize the markov model with states
	mkv, err := model.New(states)
	if err != nil {
		log.Fatalf("unable to create model. reason: %q", err)
	}
	// train the model on the training data
	mkv.Train(trainingData)
	// generate the music
	// visualize the music
}
