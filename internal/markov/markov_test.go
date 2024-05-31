package model

import (
	"reflect"
	"testing"

	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

func TestNewMarkovModel(t *testing.T) {

	states := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
	}

	got, err := New(states)
	if err != nil {
		t.Errorf("error on model creation: %q", err.Error())
	}

	want := Markov{
		States:              states,
		InitialProbabilites: []float64{0, 0},
		TransitionMatrix:    [][]float64{{0, 0}, {0, 0}},
		stateIndexes: map[notes.Note]int{
			{PitchWithOctave: "E4", QuarterNoteDuration: 1}: 0,
			{PitchWithOctave: "D4", QuarterNoteDuration: 1}: 1,
		},
	}

	if !reflect.DeepEqual(want, *got) {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}

func TestTrain(t *testing.T) {

	states := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
	}

	mkv, err := New(states)
	if err != nil {
		t.Errorf("error on model creation: %q", err.Error())
	}

	td := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
	}

	mkv.Train(td)

	got := mkv.InitialProbabilites
	want := []float64{0.5, 0.5}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("InitialProbablities don't match. WANT: %v, GOT: %v\n", want, got)
	}
}

func TestCalculateTransitionProbablities(t *testing.T) {

	states := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
	}

	mkv, err := New(states)
	if err != nil {
		t.Errorf("error on model creation: %q", err.Error())
	}

	td := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
	}

	mkv.Train(td)

	got := mkv.TransitionMatrix
	want := [][]float64{
		{0.5, 0.5, 0},
		{0, 0, 1},
		{0, 0, 0},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("transition probabilities don't match. WANT: %v, GOT: %v\n", want, got)
	}
}

func TestGenerateFirstMelody(t *testing.T) {

	states := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
	}

	mkv, err := New(states)
	if err != nil {
		t.Errorf("error on model creation: %q", err.Error())
	}

	td := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
	}

	mkv.Train(td)

	generatedMelody := mkv.Generate(1)

	if len(generatedMelody) != 1 {
		t.Error("generate inital melody didn't work\n")
	}
	t.Logf("generated initial melody of %v\n", generatedMelody)
}

func TestGenerate(t *testing.T) {

	var MELODY_LENGTH = 4

	states := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
	}

	mkv, err := New(states)
	if err != nil {
		t.Errorf("error on model creation: %q", err.Error())
	}

	td := []notes.Note{
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
		{PitchWithOctave: "D4", QuarterNoteDuration: 1},
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
	}

	mkv.Train(td)

	generatedMelody := mkv.Generate(MELODY_LENGTH)

	if len(generatedMelody) != MELODY_LENGTH {
		t.Errorf("generated melody length not equal to requested length. want: %d, got: %d", MELODY_LENGTH, len(generatedMelody))
	}

	t.Logf("generated melody: %+v\n", generatedMelody)
}
