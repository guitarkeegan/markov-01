package lilypond

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	notes "github.com/guitarkeegan/markov-go/internal/notation"
)

type Lilypond struct {
	Score          string
	Title          string
	Composer       string
	Notes          []notes.Note
	LilypondMelody string
}

func New(title, composer string, notes []notes.Note) *Lilypond {

	return &Lilypond{
		Score:          "",
		Title:          title,
		Composer:       composer,
		Notes:          notes,
		LilypondMelody: "",
	}
}

func (lp *Lilypond) writeScoreToFile() {
	os.WriteFile("tmp_scores/score.ly", []byte(lp.Score), 0666)
}

func (lp *Lilypond) DisplayScore() error {

	lp.LilypondMelody = lp.convertNotesToLilypond()
	scoreTemplate, err := template.New("score").Parse(`\version "2.24.3"
\header {
  title = "{{.Title}}"
  composer = "{{.Composer}}"
}

\score {
  \new Staff {
    \relative c' {
      \time 4/4
      \key c \major

	{{.LilypondMelody}}
    }
  }
}`)

	if err != nil {
		return errors.New("error parsing template: " + err.Error())
	}

	var scoreBuilder strings.Builder
	err = scoreTemplate.Execute(&scoreBuilder, lp)
	if err != nil {
		return errors.New("error executing template: " + err.Error())
	}

	lp.Score = scoreBuilder.String()

	lp.writeScoreToFile()

	// TODO: better way to get path?
	cmd := exec.Command("lilypond", "-fpng", "-o", "output", "tmp_scores/score.ly")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("call to lilypond failed: " + err.Error())
	}

	fmt.Printf("%s", output)

	return nil

}

func (lp *Lilypond) convertNotesToLilypond() string {

	var lilypondMelody []string

	for _, note := range lp.Notes {
		// TODO: Ignoring octave for now
		// TODO: also the string is gross
		newString := fmt.Sprintf("%s%d", strings.ToLower(string(note.PitchWithOctave[0])), convertDuration(note.QuarterNoteDuration))
		lilypondMelody = append(lilypondMelody, newString)
	}
	return strings.Join(lilypondMelody, " ")
}

// TODO: make this a float so that 8th, 16th note durations are possible.
func convertDuration(quarterNoteDuration int) int {
	return int(4 / quarterNoteDuration)
}
