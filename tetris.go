package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	cursor "github.com/ArturC03/shellcade/functions"
	"github.com/eiannone/keyboard"
)

const (
	N = 20 // Size of the arena
)

var (
	arena         [N][N]string
	width, height int // These will be set later dynamically
)

// Definição das cores das peças
var pieceColors = map[string]string{
	"I": "\033[36m", // Ciano
	"J": "\033[34m", // Azul
	"L": "\033[33m", // Amarelo
	"O": "\033[32m", // Verde
	"S": "\033[31m", // Vermelho
	"T": "\033[35m", // Magenta
	"Z": "\033[30m", // Preto
}

// Define as peças do Tetris com seus respectivos blocos
var pieces = [][][]string{
	// I
	{
		{" ", " ", " ", " "},
		{"I", "I", "I", "I"},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
	},
	// J
	{
		{" ", " ", "J"},
		{"J", "J", "J"},
		{" ", " ", " "},
	},
	// L
	{
		{"L", " ", " "},
		{"L", "L", "L"},
		{" ", " ", " "},
	},
	// O
	{
		{" ", "O", "O"},
		{" ", "O", "O"},
		{" ", " ", " "},
	},
	// S
	{
		{" ", "S", "S"},
		{"S", "S", " "},
		{" ", " ", " "},
	},
	// T
	{
		{" ", "T", " "},
		{"T", "T", "T"},
		{" ", " ", " "},
	},
	// Z
	{
		{"Z", "Z", " "},
		{" ", "Z", "Z"},
		{" ", " ", " "},
	},
}

// Função para criar uma nova peça aleatória
func createPiece() [][]string {
	rand.Seed(time.Now().UnixNano())
	return pieces[rand.Intn(len(pieces))]
}

// Função para girar a peça no sentido horário
func rotatePieceClockwise(piece [][]string) [][]string {
	size := len(piece)
	rotated := make([][]string, size)
	for i := 0; i < size; i++ {
		rotated[i] = make([]string, size)
		for j := 0; j < size; j++ {
			rotated[i][j] = piece[size-1-j][i]
		}
	}
	return rotated
}

// Função para mostrar a arena
func ShowArena() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(arena[i][j])
		}
		fmt.Println()
	}
}

// Função principal do jogo Tetris
func RunTetris() {
	cursor.Clear()

	// Limpa a arena antes de iniciar
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			arena[i][j] = " "
		}
	}

	// Mostra a arena inicial
	ShowArena()

	// Cria a primeira peça aleatória
	currentPiece := createPiece()
	posX, posY := (N-len(currentPiece))/2, 0

	// Set up keyboard input handling
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// Loop principal do jogo
	for {
		// Verifica se a peça pode continuar descendo
		if !checkCollision(currentPiece, posX, posY+1) {
			posY++
		} else {
			// Fixar a peça atual na arena
			fixPieceToArena(currentPiece, posX, posY)

			// Verifica se o jogo acabou (quando a próxima peça colide ao ser criada)
			if checkCollision(createPiece(), (N-len(currentPiece))/2, 0) {
				fmt.Println("Game Over!")
				return
			}

			// Cria uma nova peça aleatória
			currentPiece = createPiece()
			posX, posY = (N-len(currentPiece))/2, 0
		}

		// Atualiza e mostra a arena com a peça atual
		UpdateArena(currentPiece, posX, posY)
		cursor.Clear()
		ShowArena()

		// Simulação do tempo de queda da peça
		time.Sleep(500 * time.Millisecond)
	}
}

// Função para verificar colisão com a arena ou outras peças
func checkCollision(piece [][]string, posX int, posY int) bool {
	for i := 0; i < len(piece); i++ {
		for j := 0; j < len(piece[i]); j++ {
			if piece[i][j] != " " {
				arenaY := posY + i
				arenaX := posX + j
				if arenaY >= N || arenaX < 0 || arenaX >= N || (arenaY >= 0 && arena[arenaY][arenaX] != " ") {
					return true
				}
			}
		}
	}
	return false
}

// Função para fixar a peça atual na arena
func fixPieceToArena(piece [][]string, posX int, posY int) {
	for i := 0; i < len(piece); i++ {
		for j := 0; j < len(piece[i]); j++ {
			if piece[i][j] != " " {
				color := pieceColors[piece[i][j]]
				arenaY := posY + i
				arenaX := posX + j
				if arenaY >= 0 && arenaY < N && arenaX >= 0 && arenaX < N {
					arena[arenaY][arenaX] = color + "█"
				}
			}
		}
	}
}

// Função para atualizar a arena com a peça atual
func UpdateArena(piece [][]string, posX int, posY int) {
	// Limpar arena
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			arena[i][j] = " "
		}
	}

	// Adicionar peça à arena
	for i := 0; i < len(piece); i++ {
		for j := 0; j < len(piece[i]); j++ {
			if piece[i][j] != " " {
				color := pieceColors[piece[i][j]]
				arenaY := posY + i
				arenaX := posX + j
				if arenaY >= 0 && arenaY < N && arenaX >= 0 && arenaX < N {
					arena[arenaY][arenaX] = color + "█"
				}
			}
		}
	}
}
