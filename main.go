package main

import (
    "strings"
    "fmt"
    "github.com/ArturC03/shellcade/functions"
    "github.com/eiannone/keyboard"
    "log"
)

const (
  COLOR_CODE = "\033[38;2;3;169;244m"
  WHITE_BACKGROUND = "\033[47m"
  DEFAULT_COLOR= "\033[0m"
)

func Menu(options []string) int {
	cursor.Clear()
	cursor.Position(0, 0)

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Use W/S to navigate, Enter to select, and Q to quit")

	selection := 0

	for {
		// Redraw menu with the current selection
		cursor.Clear()
		cursor.Position(0, 0)
		for i := 0; i < len(options); i++ {
			if i == selection {
				fmt.Printf(COLOR_CODE + options[i] + DEFAULT_COLOR + "\n")
			} else {
				fmt.Println(options[i])
			}
		}

		// Read key
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
    
    if key == keyboard.KeyEnter {
			fmt.Printf("You selected option %d\n", selection)
			return selection
    }

		switch char {
		case 'q':
			fmt.Println("Exiting")
			return -1
		case 'w':
			if selection > 0 {
				selection--
			}
		case 's':
			if selection < len(options)-1 {
				selection++
			}
		case '\n': // Enter key
			fmt.Printf("You selected option %d\n", selection)
			return selection
		}
	}
}

func StartScreen(width, height int)  {


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
--------------------------------------------------------------------------
                         ~ CLICK ENTER TO START ~                         
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
  fmt.Print(DEFAULT_COLOR) 

    // Set cursor position and print ASCII art
    fmt.Scanln()
    cursor.Clear()
}

func main() {
    
  width, height, _:= cursor.GetScreenSize()

  //Show StartScreen
  StartScreen(width, height)
  
  // Menu()
  options := []string{"Option 1", "Option 2", "Option 3"}
  Menu(options)
  cursor.Show()
}

