package greeting

import "fmt"

// HelloWorld function that return a string value for greeting.
func HelloWorld() string {

	return "Hello, World!"
}

//main function to call the HelloWorld() function.
func main() {

	//printing to the screen
	fmt.Print(HelloWorld())
}
