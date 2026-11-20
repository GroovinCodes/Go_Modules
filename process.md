// - Part 1
// 1. Create different go files and modules with their respective modules
// 2. Import the age pacakge, and provide the values for the function within the age file from the handler func
// 3. In the terminal, provide the following command(s) to update the go module of greeting.go
// go mod edit -replace Go_Modules/age=../age
// go mod tidy
// Then run the main file

// Because we're not going to publish the package for go to download from the internet, we use the
// command above to make go look for the package locally

// - Part 2
// Then using the log and errors package, we implement error handling to stop the program in the case
// That we come across an empty input

// - Part 3
// Create a slice of greetings that randomly selects from the slice
// In the age program, simply create a function that returns a string taken from a slice of them. 
// Each string is a sentence that matches with the value of name and age, being selected using the math/rand package 
// via their index. Once done, the function simply returns the statement, Details returns the message back to the 
// age var withing greeting.go, and prints out the output. 

// - Part 4
// We now use the slice to greet multiple people instead of just one!
// Since changing the original function within the age program would change the func signature 
// and ruin programs, we'll run a different function that handles that case for backwards compatibility dubbed "multipleGreetings"

// The process is simple: We create a slice of names (and ages since you're using them) within greeting and pass them into age to be put within the function.
// Now I changed the program to be a bit more specific to our situation. Instead of using key-value pairs, I decided to two functions provide random values 
// to act as some sort of random sentence generator through the use of for loops.
// I also wanted to make every operation have its own individual placement for function efficiency and clear understanding.
// For future projects of trying to see if a slice contains something, use this over a for loop:
if slices.Contains(age, "") {
			return "", errors.New("Empty value detected")
		}

// - Part 5
// We will now add tests into our project to test for bugs and problems 



Aftermath Project Details:
// Have the program choose random names and ages from a slice to each statement unique (Done - 11/20/2026)
// Implement a switch-case statement menu structure that is controllable from the terminal while implementing user input 