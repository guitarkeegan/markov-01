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
		stateIndexes:        map[interface{}]int{},
	}

	if !reflect.DeepEqual(want, *got) {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}
