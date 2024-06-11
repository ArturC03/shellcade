package main

import "fmt"
import (
    "github.com/ArturC03/shellcade/functions"
)

func main() {
    
  width, height, _:= cursor.GetScreenSize()
  cursor.Clear()

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
// Calculate starting x-coordinate for centering horizontally
    startX := width/2 - 37

    // Hide cursor
    cursor.Hide()

    // Set cursor position and print ASCII art
    cursor.Position(startX, height/2-4)
    fmt.Println(asciiArt)
    fmt.Scanln()
    cursor.Clear()

    
}

