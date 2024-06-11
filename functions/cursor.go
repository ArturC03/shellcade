package cursor

import "fmt" 
// import "errors"


func Hide() {
  fmt.Print("\033[?25l")
}

func MoveU()  {
  
}

func MoveD()  {
  
}

func MoveL() {
  
}

func MoveR(n int) error {
  
  if (0 > n) {
     n *= -1
  }

    fmt.Printf("\033[%dC", n)
    return nil
}

func Position()  {
  
}

func Show() {
  fmt.Print("\033[?25h")
}

