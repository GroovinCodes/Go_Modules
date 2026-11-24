package main

import (
	"Go_Modules/age"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("age: ")
	log.SetFlags(0)

	names := []string{"Gabe", "Sarah", "Michael"}
	ages := []string{"20", "24", "81"}

	for i := 0; i < len(names); i++ { // Used to output a custom sentence
		message, err := age.Details(ages, names) // relates to what values will receive from age, not what will be passed to it. What passes is only our parameters
		if err != nil {
			log.Fatal(err)
		}
		// err will receive the value of either the errors.New() or nil, depending on the Details func within age. That is
		// returned, while what is passed is our string values. The age var is just to take in the print statement that is returned.

		fmt.Println(message)
	}

}
