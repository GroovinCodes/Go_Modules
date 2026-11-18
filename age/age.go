package age

import (
	"errors"
	"fmt"
	"math/rand"
)

func Details(age string, name string) (string, error) {

	if age == "" || name == "" {
		return "", errors.New("Values returned were empty")
	}

	message := fmt.Sprintf(randomSentence(), name, age)
	return message, nil
}

func randomSentence() string {
	// create a slice full of sentences that will be randomly chosen
	sentence := [...]string{
		"Hey, %v! I found out that you're %v years old now!",
		"Woah %v, when'd you become %v?",
		"%v just turned %v last week.",
	}

	// return a random sentence via a random index
	return sentence[rand.Intn(len(sentence))]
}
