package model

import (
	"errors"
	"log"
	"math/rand"

	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

type Markov struct {
	States              []notes.Note
	InitialProbabilites []float64
	TransitionMatrix    [][]float64
	stateIndexes        map[notes.Note]int
}

type MarkovChain interface {
	Train(n []notes.Note)
	Generate(l int) []notes.Note
}

func (m *Markov) Train(n []notes.Note) {

	m.calculateInitialProbablities(n)
	m.calculateTransitionMatrix(n)
}

func (m *Markov) calculateInitialProbablities(n []notes.Note) {

	var rowSum float64

	for _, note := range n {
		idx := m.stateIndexes[note]
		m.InitialProbabilites[idx]++
		rowSum++
	}
	// normalize row
	for i, score := range m.InitialProbabilites {
		if score != 0 {
			m.InitialProbabilites[i] = score / rowSum
		}
	}
}

func (m *Markov) calculateTransitionMatrix(n []notes.Note) {
	// increment transitions
	for i := range n[:len(n)-1] {
		idx := m.stateIndexes[n[i]]
		nextIdx := m.stateIndexes[n[i+1]]
		m.TransitionMatrix[idx][nextIdx]++
	}
	for i, row := range m.TransitionMatrix {
		// TODO: normalize rows
		m.TransitionMatrix[i] = m.normalizeMatrix(row)
	}
}

func (m *Markov) normalizeMatrix(row []float64) []float64 {
	// sum the row
	var total float64
	for _, val := range row {
		total += val
	}
	// check for zeros
	if total == 0 {
		return row
	}

	for i, val := range row {
		row[i] = val / total
	}
	return row

}

func (m *Markov) Generate(length int) []notes.Note {
	// generate a melody of a given length

	// get the first note
	// it seems like at compile time, the InitialProbabilites are nil
	firstNote, err := m.generateFirstNote()
	if err != nil {
		log.Fatalf("%q", err)
	}

	melody := []notes.Note{firstNote}
	// generate the next note, based on the last note, using the probablity matrix
	for i := 1; i < length; i++ {
		nextNote, err := m.generateNextNote(melody[len(melody)-1])
		if err != nil {
			log.Fatalf("unable to generate next note: %q", err)
		}
		melody = append(melody, nextNote)
	}
	return melody
}

// this is lame
func (m *Markov) generateFirstNote() (notes.Note, error) {
	randomNumber := rand.Float64()

	cumulativeProbablities := make([]float64, len(m.InitialProbabilites))
	cumulativeProbablities[0] = m.InitialProbabilites[0]
	for i := 1; i < len(m.InitialProbabilites); i++ {
		cumulativeProbablities[i] = cumulativeProbablities[i-1] + m.InitialProbabilites[i]
	}

	for i, val := range cumulativeProbablities {
		if randomNumber < val {
			for k, v := range m.stateIndexes {
				if i == v {
					return k, nil
				}
			}
		}
	}

	return notes.Note{
		PitchWithOctave:     "",
		QuarterNoteDuration: 0,
	}, errors.New("couldn't generate first note")

}

func (m *Markov) generateNextNote(n notes.Note) (notes.Note, error) {

	// is there a possible next note?
	if m.hasSubsequestState(m.stateIndexes[n]) {

		randomNumber := rand.Float64()

		cumulativeProbablities := make([]float64, len(m.InitialProbabilites))
		cumulativeProbablities[0] = m.TransitionMatrix[m.stateIndexes[n]][0]
		for i := 1; i < len(m.InitialProbabilites); i++ {
			cumulativeProbablities[i] = cumulativeProbablities[i-1] + m.TransitionMatrix[m.stateIndexes[n]][i]
		}

		for i, val := range cumulativeProbablities {
			if randomNumber < val {
				for k, v := range m.stateIndexes {
					if i == v {
						return k, nil
					}
				}
			}
		}
		return notes.Note{}, errors.New("unable to generateNextNote")

	} else {
		return m.generateFirstNote()

	}
}

func (m *Markov) hasSubsequestState(currentState int) bool {

	row := m.TransitionMatrix[currentState]
	var total float64
	for _, num := range row {
		total += num
	}
	if total > 0 {
		return true
	}
	return false
}

func New(states []notes.Note) (*Markov, error) {

	var (
		MIN_STATES = 1
		MAX_STATES = 30
	)
	// TODO: this was an effor to fix the Out of bounds on generate
	if len(states) < MIN_STATES || len(states) >= MAX_STATES {
		return nil, errors.New("need between 1 - 30")
	}

	mkv := &Markov{
		States:              states,
		InitialProbabilites: make([]float64, len(states)),
		TransitionMatrix:    make([][]float64, len(states)),
		stateIndexes:        map[notes.Note]int{},
	}

	// create the transition matrix
	for i := range states {
		mkv.TransitionMatrix[i] = make([]float64, len(states))
		mkv.stateIndexes[states[i]] = i
	}

	return mkv, nil
}
