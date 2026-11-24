package age

// all we're doing is basically what greeting does, just with tests.
import (
	"regexp"
	"strings"
	"testing"
)

// we'll be running to see if details receives its proper values of a slice of strings.
func TestDetailsValues(t *testing.T) {
	// create age and name
	names := []string{"Rafy", "Mikey", "Thomas"}
	ages := []string{"10", "24", "38"}

	// Make slices into strings. This is to have the program still be able to randomly choose from whatever's in the slice,
	// but also compile the sentence as a single string
	wantNames := regexp.MustCompile(`\b(` + strings.Join(names, "|") + `)\b`)
	wantAges := regexp.MustCompile(`\b(` + strings.Join(ages, "|") + `)\b`)

	// pass it all to Details, and create cases if what was returned doesn't match with what was wanted
	message, err := Details(ages, names)
	if err != nil {
		t.Fatalf("Details had an unexpected error: %v", err)
	}

	if !wantNames.MatchString(message) {
		t.Errorf("Message %q contained none of the values %v", message, names)
	}

	if !wantAges.MatchString(message) {
		t.Errorf("Message %q contained none of the values %v", message, ages)
	}
}

// Now, create a function that tests if the error handling of age is working properly.
func TestEmptyValues(t *testing.T) {
	message, err := Details([]string{""}, []string{""}) // Since Details takes in slices, pass slices with an empty string to activate the error handling
	if message != "" || err == nil {                    //  Since the default value of strings in go is "", it'll return the error in age and make this function pass
		t.Errorf("Error handling failed.")
	}
}

// Optional: Provide a for loop of how many times you want the operation to be ran for more possibilities of finding a fail. Like iterating 50 times.
