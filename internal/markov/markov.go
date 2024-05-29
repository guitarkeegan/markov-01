package model

import (
	"errors"

	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

type Markov struct {
	States              []notes.Note
	InitialProbabilites []float64
	TransitionMatrix    [][]float64
	stateIndexes        map[interface{}]int
}

func New(states []notes.Note) (*Markov, error) {

	if len(states) < 1 {
		return nil, errors.New("need at least 1 initial state")
	}

	mkv := &Markov{
		States:              states,
		InitialProbabilites: make([]float64, len(states)),
		TransitionMatrix:    make([][]float64, len(states)),
		stateIndexes:        map[interface{}]int{},
	}

	// create the transition matrix
	for i := range states {
		mkv.TransitionMatrix[i] = make([]float64, len(states))
	}

	return mkv, nil
}
