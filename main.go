package main

import "fmt"
import "github.com/ArturC03/shellcade/functions"

func main() {
    fmt.Print("\033[?25h")
    asciiArt := `
  ██████  ██░ ██ ▓█████  ██▓     ██▓     ▄████▄   ▄▄▄      ▓█████▄ ▓█████ 
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    ▒██▀ ▀█  ▒████▄    ▒██▀ ██▌▓█   ▀ 
░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    ▒▓█    ▄ ▒██  ▀█▄  ░██   █▌▒███   
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░    ▒▓▓▄ ▄██▒░██▄▄▄▄██ ░▓█▄   ▌▒▓█  ▄ 
▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒▒ ▓███▀ ░ ▓█   ▓██▒░▒████▓ ░▒████▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░░ ░▒ ▒  ░ ▒▒   ▓▒█░ ▒▒▓  ▒ ░░ ▒░ ░
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░  ░  ▒     ▒   ▒▒ ░ ░ ▒  ▒  ░ ░  ░
░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░   ░          ░   ▒    ░ ░  ░    ░   
      ░   ░  ░  ░   ░  ░    ░  ░    ░  ░░ ░            ░  ░   ░       ░  ░
`
    cursor.Hide()
    fmt.Println(asciiArt)
    fmt.Scanln()
}

