package main

import "fmt"

type Game struct {
	player1 *Player
	player2 *Player
}

func (g *Game) start() {
	// start initiates and runs a game between two players.
	// Each player takes turns attacking the other until one of them dies.
	// It prints a message at the start of the game and a message when the game ends.

	fmt.Println("Starting the game...")

	// Continue the game loop until one of the players is dead.
	for !g.player1.isDead() && !g.player2.isDead() {
		// Determine which player should attack based on their current health.
		if g.player1.health <= g.player2.health {
			// Player 1 attacks Player 2.
			g.player1.attackPlayer(g.player2)
		} else {
			// Player 2 attacks Player 1.
			g.player2.attackPlayer(g.player1)
		}

		// Check if Player 1 has died and declare Player 2 as the winner.
		if g.player1.isDead() {
			fmt.Printf("%s has died. %s wins!\n", g.player1.name, g.player2.name)
			return
		} else if g.player2.isDead() {
			// Check if Player 2 has died and declare Player 1 as the winner.
			fmt.Printf("%s has died. %s wins!\n", g.player2.name, g.player1.name)
			return
		}
	}
}
