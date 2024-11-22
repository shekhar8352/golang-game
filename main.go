package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	// main is the entry point for the game.
	// It seeds the random number generator, initializes two players with different attributes,
	// displays their initial attributes, and then starts the game.

	// Seed the random number generator for dice rolls
	rand.Seed(time.Now().UnixNano())

	// Initialize players with their respective attributes
	player1 := &Player{name: "Player 1", health: 50, strength: 5, attack: 10}
	player2 := &Player{name: "Player 2", health: 100, strength: 10, attack: 5}

	// Display initial player attributes
	fmt.Printf("Initial stats - %s: health=%d, strength=%d, attack=%d\n", player1.name, player1.health, player1.strength, player1.attack)
	fmt.Printf("Initial stats - %s: health=%d, strength=%d, attack=%d\n", player2.name, player2.health, player2.strength, player2.attack)
	fmt.Println()

	// Start the game
	game := Game{player1: player1, player2: player2}
	game.start()
}
