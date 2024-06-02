package lilypond

import (
	"testing"

	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

func TestConvertNoteToLilypond(t *testing.T) {

	notes := []notes.Note{
		{PitchWithOctave: "C4", QuarterNoteDuration: 1},
		{PitchWithOctave: "E4", QuarterNoteDuration: 1},
	}

	lp := New("Test Title", "Test Composer", notes)

	got := lp.convertNotesToLilypond()

	want := "C4 E4"

	if got != want {
		t.Errorf("conversion to lilypond string no good, WANT: %s\nGOT: %s", want, got)
	}
}
