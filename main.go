package main

import (
    "strings"
    "fmt"
    "github.com/ArturC03/shellcade/functions"
)

func StartScreen(width, height int)  {

  // Set the text color to #03A9F4 (true color)
  COLOR_CODE:= "\033[38;2;3;169;244m"

  asciiArt := 
`  ██████  ██░ ██ ▓█████  ██▓     ██▓     ▄████▄   ▄▄▄      ▓█████▄ ▓█████
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    ▒██▀ ▀█  ▒████▄    ▒██▀ ██▌▓█   ▀ 
░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    ▒▓█    ▄ ▒██  ▀█▄  ░██   █▌▒███   
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░    ▒▓▓▄ ▄██▒░██▄▄▄▄██ ░▓█▄   ▌▒▓█  ▄ 
▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒▒ ▓███▀ ░ ▓█   ▓██▒░▒████▓ ░▒████▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░░ ░▒ ▒  ░ ▒▒   ▓▒█░ ▒▒▓  ▒ ░░ ▒░ ░
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░  ░  ▒     ▒   ▒▒ ░ ░ ▒  ▒  ░ ░  ░
░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░   ░          ░   ▒    ░ ░  ░    ░   
      ░   ░  ░  ░   ░  ░    ░  ░    ░  ░░ ░            ░  ░   ░       ░  ░
`
  // Clear Screen
  cursor.Clear()
  // Hide cursor
  cursor.Hide()

  lines := strings.Split(asciiArt, "\n")

  fmt.Print(COLOR_CODE)
  
  // Calculate the starting y-coordinate for centering vertically
	startY := height/2 - 4
	
  for i, line := range lines {
		// Calculate the x position to center the line
		startX := width/2 - 74/2
		cursor.Position(startX, startY+i)
		fmt.Print(line)
	}
  DEFAULT_COLOR := "\033[0m"
  fmt.Print(DEFAULT_COLOR)
    // Set cursor position and print ASCII art
    fmt.Scanln()
    cursor.Clear()
}

func main() {
    
  width, height, _:= cursor.GetScreenSize()

  //Show StartScreen
  StartScreen(width, height)
}

