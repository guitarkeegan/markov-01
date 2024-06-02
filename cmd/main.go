package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/guitarkeegan/markov-go/internal/lilypond"
	model "github.com/guitarkeegan/markov-go/internal/markov"
	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

var DEFAULT_MELODY_LENGTH = 4

func main() {

	var melodyLength int
	// get # of notes form user input
	flag.IntVar(&melodyLength, "l", 4, "desired melody length to generate")
	flag.Parse()

	// create training data
	trainingData := createTrainingData()
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
	melody := mkv.Generate(melodyLength)
	// visualize the music
	// TODO: call out to lilypond
	lp := lilypond.New("K's title", "Keegan", melody)
	lp.DisplayScore()

	fmt.Printf("%v\n", melody)
}
