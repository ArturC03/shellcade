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
		"█", "▒", "░", "▓", "░", "▒", "▓", "░",
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
	directionCh := make(chan string, 1) // Buffered channel for direction input

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
					directionCh <- "up"
				}
			case 's':
				if direction != "up" {
					directionCh <- "down"
				}
			case 'a':
				if direction != "right" {
					directionCh <- "left"
				}
			case 'd':
				if direction != "left" {
					directionCh <- "right"
				}
			}
		}
	}()

	// Function to check if the point is within the screen boundaries
	isWithinBounds := func(p Point) bool {
		width, height, _ := cursor.GetScreenSize() // Corrected: Get screen size here
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
			// Check direction input from channel
			select {
			case dir := <-directionCh:
				direction = dir
			default:
				// No new direction input, continue with current direction
			}

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
				return
			}

			// Add new head to the snake
			snake = append([]Point{newHead}, snake...)

			// Check for self-collision
			if hasSelfCollision(snake) {
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
				letter := snakeLetter[i]
				fmt.Print(SNAKE_COLOR + letter + DEFAULT_COLOR)
				i++
			}

			// Print the fruit
			cursor.Position(fruit.x, fruit.y)
			fmt.Print(FRUIT_COLOR + fruitLetter + DEFAULT_COLOR)

			// Delay for a while to control the game speed
			time.Sleep(120 * time.Millisecond)
		}
	}
}


