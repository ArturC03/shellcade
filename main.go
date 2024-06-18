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

func Menu(options []string, biggestString string, width int, height int) int {
	cursor.Clear()
	cursor.Position((width - len(biggestString))/2,(height - len(options))/2)

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	selection := 0

	for {
		// Redraw menu with the current selection
		cursor.Clear()
  	cursor.MakeOutline((width - len(biggestString))/2 - 1,(height - len(options))/2 - 1, len(biggestString)+2, len(options) + 2)
		cursor.Position((width - len(biggestString))/2,(height - len(options))/2)
		for i := 0; i < len(options); i++ {
			if i == selection {
				cursor.Position((width - len(biggestString))/2,(height - len(options))/2 + i)
				fmt.Printf(COLOR_CODE + options[i] + DEFAULT_COLOR + "\n")
			} else {
				cursor.Position((width - len(biggestString))/2,(height - len(options))/2 + i)
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
			keyboard.Close()
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
	
  // Calculate the starting x-coordinate for centering vertically
	startX := width/2 - 74/2
  for i, line := range lines {
		// Calculate the x position to center the line
		cursor.Position(startX, startY+i)
		fmt.Print(line)
	}
  cursor.MakeOutline(startX - 1,startY - 1,74 + 3, 8 + 2)
  fmt.Print(DEFAULT_COLOR) 

  if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

  for { // Read key
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
    
    if key == keyboard.KeyEnter {
      break
    }
  }

    keyboard.Close()
    cursor.Clear()
}

func main() {

  fmt.Scanln()
  
  width, height, _ := cursor.GetScreenSize()

  StartScreen(width, height)

  options := []string{"1. Snake Game", "2. Tetris", "3. Shell Fight"}
  biggestString := options[0]
  for i := 1; i < len(options);i++ {
    if biggestString < options[i] {
      biggestString = options[i]
    }
  }

	option := Menu(options, biggestString, width, height)
	if (0 == option) {
		RunSnake()
	}

	cursor.Show()

}
