/* The package “main” tells the Go compiler that the package should compile as an executable program 
 * instead of a shared library. */
package main

/* Package fmt implements formatted I/O with functions analogous to C's printf and scanf */
import "fmt"

/* The main function in the package “main” will be the entry point of our executable program. */
func main() {
	/* he fmt.Println() function in Go language formats using the default formats for its operands 
	 * and writes to standard output. */
	fmt.Println("Hello, world.")
}

/* Execute the program by writing in terminal $ go run Hello.go */
