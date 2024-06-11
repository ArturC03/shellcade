package cursor// package declaration

import "fmt" // import necessary packages

// MyFunction is a custom function
func Hide() {
  fmt.Print("\033[?25l")
}

func Show() {
  fmt.Print("\033[?25h")
}

