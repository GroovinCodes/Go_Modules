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
