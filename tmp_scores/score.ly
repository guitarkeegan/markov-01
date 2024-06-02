\version "2.24.3"
\header {
  title = "K's title"
  composer = "Keegan"
}

\score {
  \new Staff {
    \relative c' {
      \time 4/4
      \key c \major

	b4 a4 g4 c4 b4 a4 e4 c4 b4 c4 g4 a4 g4 c4 b4 c4
    }
  }
}