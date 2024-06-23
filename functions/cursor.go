package cursor

import "fmt" 

// MoveR moves the cursor to the right by n spaces
func MoveR(n int) {
    if n < 0 {
        n *= -1 // turn n into a positive integer
    }
    // ANSI escape code to move the cursor to the right by n spaces
    fmt.Printf("\033[%dC", n)
}

// MoveL moves the cursor to the left by n spaces
func MoveL(n int) {
    if n < 0 {
        n *= -1 // turn n into a positive integer
    }
    // ANSI escape code to move the cursor to the left by n spaces
    fmt.Printf("\033[%dD", n)
}

// MoveD moves the cursor down by n lines
func MoveD(n int) {
    if n < 0 {
        n *= -1 // turn n into a positive integer
    }
    // ANSI escape code to move the cursor down by n lines
    fmt.Printf("\033[%dB", n)
}

// MoveU moves the cursor up by n lines
func MoveU(n int) {
    if n < 0 {
        n *= -1 // turn n into a positive integer
    }
    // ANSI escape code to move the cursor up by n lines
    fmt.Printf("\033[%dA", n)
}

// Position changes the cursor position
func Position(x, y int) {
    // ANSI escape code to set cursor position
    fmt.Printf("\033[%d;%df", y, x)
}

// Hide hides the cursor
func Hide() {
  fmt.Print("\033[?25l")
}

func Show() {
  fmt.Print("\033[?25h")
}

