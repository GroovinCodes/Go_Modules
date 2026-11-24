package age

import (
	"errors"
	"fmt"
	"math/rand"
)

func Details(age []string, name []string) (string, error) { // edit to take in slices and return a message via a for loop
	// To check if the slices contain empty values, you must iterate through them. (There is a cleaner way of doing this)
	for _, ages := range age {
		if ages == "" {
			return "", errors.New("Empty value detected")
		}
	}

	for _, names := range name {
		if names == "" {
			return "", errors.New("Empty value detected")
		}
	}
	// after this error handling, we can pass the slices into multipleGreetings to be added to the statements

	message := multipleGreetings(age, name) // What will be returned should be done multiple times to account for all values within the slice.
	return message, nil
}

// Instead of using key value pairs, why not try and provide two slices of names to assign as parameters to the Details function.
// this func takes in the slices, calls the randomSentence function to get a sentence, and random chooses from the slices to form a sentence and return it back to
// Details.
func multipleGreetings(age []string, name []string) string { // returns a string for message. Needs to select an element from age and name by random
	nameIndex := rand.Intn(len(name)) // Have the program choose the random index, then passed the value at that index back to the function
	ageIndex := rand.Intn(len(age))

	randomName := name[nameIndex]
	randomAge := age[ageIndex]

	message := fmt.Sprintf(randomSentence(), randomName, randomAge)
	return message
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

// Optional: Add error handling in the case that the length of passed slices are less than 0
