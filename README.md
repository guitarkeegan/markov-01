# Markov Melody Generator

## What is this?

This is an implementation of a Markov chain that has been trained on the first phrase of
"St. Thomas". 

## Dependancies

You will need to have lilypond installed on your computer for the score output to work
properly. 

## How to run

1. Clone the project
2. From the root ```go download```
3. Run the program ```go run ./cmd```
4. Optional - You can also pass a length to make the generated melody longer ```go run ./cmd -l 12```

## What else?

You can experiment with changing the training set in the cmd/training.go file.

### TODOs

In no particular order

- update int to float on note durations
- get user input for title and composer

### Contribute

If you'd like to work on the project with me, just send me an [email](keegananglim@gmail.com).
