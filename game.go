package main

import "fmt"

type Game struct {
	player1 *Player
	player2 *Player
}

func (g *Game) start() {
	// start initiates and runs a game between two players.
	// Each player takes turns attacking the other until one of them dies.

	fmt.Println("Starting the game...")
	
	// Determine which player attacks first based on their health.
	// The player with lower health should attack first.
	player1Attacks := g.player1.health <= g.player2.health
	if player1Attacks {
		g.player1.attackPlayer(g.player2)
		g.player2.attackPlayer(g.player1)
	} else {
		g.player2.attackPlayer(g.player1)
		g.player1.attackPlayer(g.player2)
	}

	// Continue the game loop until one of the players is dead.
	for !g.player1.isDead() && !g.player2.isDead() {
		if player1Attacks {
			g.player1.attackPlayer(g.player2)
			g.player2.attackPlayer(g.player1)
		} else {
			g.player2.attackPlayer(g.player1)
			g.player1.attackPlayer(g.player2)
		}

		// Check if any player has died and if so, print the winner.
		if g.player1.isDead() {
			fmt.Printf("%s has died. %s wins!\n", g.player1.name, g.player2.name)
		} else if g.player2.isDead() {
			fmt.Printf("%s has died. %s wins!\n", g.player2.name, g.player1.name)
		}
	}
}
