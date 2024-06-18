package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ArturC03/shellcade/functions"
	"github.com/eiannone/keyboard"
)

type Point struct {
	x, y int
}

const (
	SNAKE_COLOR = "\033[32m"
	FRUIT_COLOR = "\033[31m"
)

func randomSnakeCaracter() string {
		snakeCharacters := []string{
		"█", "▒", "░", "▓", "▄", "▀", "╬", "¤", "■", "▌", "▐", "▬", "≡", "☻", "◙", "⊕", "⌐", "╕", "╣", "╗", "░", "⌠", "╩", "≈", "┘", "┐", "∟", "└", "┴", "╨", "Ω", "♣", "♠", "♦", "♫", "☼", "►", "↕", "¶", "☺", "♪", "∞", "♀", "▀", "♂", "┼", "╪", "▒", "▓", "░", "├", "≥", "≤", "µ", "√", "∫", "℅", "§", "↑", "↓", "←", "→", "∟", "↔", "▲", "▼", "!", "¬", "#", "£", "¤", "%", "&", "*", "(", ")", "-", "+", "=", "@", "©", "®", "€", "™", "¥", "≠", "÷", "×", "∆", "Δ", "◊", "æ", "œ", "ø", "Ø", "∏", "π", "∑", "∂", "∫", "√", "≈", "∞", "°", "÷", "µ", "Ω", "¤", "Ω", "♠", "♣", "♦", "♫", "♪", "♥", "♂", "♀", "♪", "♫", "☼", "►", "↕", "¶", "§", "▬", "↨", "↑", "↓", "←", "→", "∟", "↔", "▲", "▼", "!", "¬", "#", "£", "¤", "%", "&", "*", "(", ")", "-", "+", "=", "@", "©", "®", "€", "™", "¥", "≠", "÷", "×", "∆", "Δ", "◊", "æ", "œ", "ø", "Ø", "∏", "π", "∑", "∂", "∫", "√", "≈", "∞", "°", "÷", "µ", "Ω", "¤", "Ω", "♠", "♣", "♦", "♫", "♪", "♥", "♂", "♀", "♪", "♫", "☼", "►", "↕", "¶", "§", "▬", "↨", "↑", "↓", "←", "→", "∟", "↔", "▲", "▼", "!", "¬", "#", "£", "¤", "%", "&", "*", "(", ")", "-", "+", "=", "@", "©", "®", "€", "™", "¥", "≠", "÷", "×", "∆", "Δ", "◊", "æ", "œ", "ø", "Ø", "∏", "π", "∑", "∂", "∫", "√", "≈", "∞", "°", "÷", "µ", "Ω",
	}
	rand.Seed(time.Now().UnixNano())
	return snakeCharacters[rand.Intn(len(snakeCharacters))]

}

func RunSnake() {
	cursor.Clear()
	width, height, _ := cursor.GetScreenSize()
	var snakeLetter []string
	snakeLetter = append(snakeLetter, randomSnakeCaracter())
	var snakeSize int = 1
	var fruitLetter string = "🍎"
	snake := []Point{{x: width / 2, y: height / 2}}

	// Function to place a new fruit at a random position
	newFruit := func() Point {
		rand.Seed(time.Now().UnixNano())
		return Point{x: rand.Intn(width-1) + 1, y: rand.Intn(height-1) + 1}
	}
	fruit := newFruit()

	direction := "down"
	quit := make(chan bool)

	// Set up keyboard input handling
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// Goroutine to handle keyboard input
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatal(err)
			}
			if key == keyboard.KeyEsc {
				quit <- true
				return
			}
			switch char {
			case 'w':
				if direction != "down" {
					direction = "up"
				}
			case 's':
				if direction != "up" {
					direction = "down"
				}
			case 'a':
				if direction != "right" {
					direction = "left"
				}
			case 'd':
				if direction != "left" {
					direction = "right"
				}
			}
		}
	}()

	// Function to check if the point is within the screen boundaries
	isWithinBounds := func(p Point) bool {
	width, height, _ = cursor.GetScreenSize()
		return p.x > 0 && p.x < width && p.y > 0 && p.y < height
	}

	// Function to check if the head collides with the body
	hasSelfCollision := func(snake []Point) bool {
		head := snake[0]
		for _, segment := range snake[1:] {
			if head == segment {
				return true
			}
		}
		return false
	}

	// Game loop
	for {
		select {
		case <-quit:
			fmt.Println("Game Over!")
			return
		default:
			// Update snake position based on direction
			head := snake[0]
			newHead := head

			switch direction {
			case "up":
				newHead.y--
			case "down":
				newHead.y++
			case "left":
				newHead.x--
			case "right":
				newHead.x++
			}

			// Check for collision with borders
			if !isWithinBounds(newHead) {
				fmt.Println("Game Over! You hit the border.")
				return
			}

			// Add new head to the snake
			snake = append([]Point{newHead}, snake...)

			// Check for self-collision
			if hasSelfCollision(snake) {
				fmt.Println("Game Over! You hit yourself.")
				return
			}

			// Check for collision with fruit
			if newHead == fruit {
				snakeLetter = append(snakeLetter, randomSnakeCaracter())
				snakeSize++
				fruit = newFruit()
			}

			// Remove the last part of the snake to simulate movement
			if len(snake) > snakeSize {
				snake = snake[:len(snake)-1]
			}

			// Clear the screen and print the snake's new position
			cursor.Clear()
			i := 0
			for _, p := range snake {
				cursor.Position(p.x, p.y)
				letter := string(snakeLetter[i]) 
				fmt.Print(SNAKE_COLOR + letter + DEFAULT_COLOR)
				i++	
			}

			// Print the fruit
			cursor.Position(fruit.x, fruit.y)
			fmt.Print(FRUIT_COLOR + fruitLetter + DEFAULT_COLOR)
			
			// Delay for a while to control the game speed
			if direction == "up" || direction == "down" {
				time.Sleep(70 * 2 * time.Millisecond)
			} else {
				time.Sleep(60* time.Millisecond)
			}
		}
	}
	keyboard.Close()
}

