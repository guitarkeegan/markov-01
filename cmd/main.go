package main

import "fmt"

func main() {

	// create training data
	trainingData := createTrainingData()
	fmt.Println(trainingData)
	// define a series of states
	states := []Note{
		Note{"C4", 1},
		Note{"D4", 1},
		Note{"E4", 1},
		Note{"F4", 1},
		Note{"G4", 1},
		Note{"A4", 1},
		Note{"B4", 1},
		Note{"C4", 2},
		Note{"D4", 2},
		Note{"E4", 2},
		Note{"F4", 2},
		Note{"G4", 2},
		Note{"A4", 2},
		Note{"B4", 2},
	}
	// initialize the markov model with states
	mkv := markov.New(states)
	// train the model on the training data
	// generate the music
	// visualize the music
}
