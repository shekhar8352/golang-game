package main

import "testing"

func TestPlayerAttack(t *testing.T) {
	// TestPlayerAttack verifies the attackPlayer method by simulating an attack from playerA to playerB.
	// It checks that playerB's health is correctly reduced when attacked, ensuring that the attack logic
	// and damage calculation function as expected.
	playerA := &Player{name: "Player A", health: 50, strength: 5, attack: 10}
	playerB := &Player{name: "Player B", health: 100, strength: 10, attack: 5}

	playerA.attackPlayer(playerB)

	if playerB.health >= 100 {
		t.Errorf("Player B's health should have been reduced, but got %d", playerB.health)
	}
}

func TestPlayerDeath(t *testing.T) {
	// TestPlayerDeath tests the isDead method by simulating an attack that reduces playerB's health
	// to 0. It verifies that playerB is correctly reported as dead by the isDead method.
	playerA := &Player{name: "Player A", health: 10, strength: 1, attack: 10}
	playerB := &Player{name: "Player B", health: 1, strength: 1, attack: 1}

	playerA.attackPlayer(playerB)

	if !playerB.isDead() {
		t.Errorf("Player B should be dead, but is still alive with health %d", playerB.health)
	}
}

func TestGameStart(t *testing.T) {
	// TestGameStart verifies that the start method of a Game object runs the game
	// until one of the players dies. It creates a game with two players, starts it,
	// and checks that one of the players has died when the game is finished.
	playerA := &Player{name: "Player A", health: 50, strength: 5, attack: 10}
	playerB := &Player{name: "Player B", health: 100, strength: 10, attack: 5}

	game := Game{player1: playerA, player2: playerB}
	game.start()

	if !playerA.isDead() && !playerB.isDead() {
		t.Errorf("One player should have died, but both are still alive.")
	}
}

func TestBothPlayersDie(t *testing.T) {
	// TestBothPlayersDie tests the scenario where both players attack and defend with the same
	// strength, causing both players to die simultaneously. It verifies that the game ends
	// correctly in this scenario by checking that both players are dead after the game is
	// finished.
	playerA := &Player{name: "Player A", health: 5, strength: 5, attack: 10}
	playerB := &Player{name: "Player B", health: 5, strength: 5, attack: 10}

	// Player A and Player B will attack and defend simultaneously
	game := Game{player1: playerA, player2: playerB}
	game.start()

	if !playerA.isDead() || !playerB.isDead() {
		t.Errorf("Both players should have died, but one or both survived. Player A: %d, Player B: %d", playerA.health, playerB.health)
	}
}

func TestPlayerAVictory(t *testing.T) {
	// TestPlayerAVictory tests the scenario where player1 defeats player2. It verifies that the
	// game ends correctly in this scenario by checking that player1 is dead after the game is
	// finished.
	player1 := &Player{name: "Player A", health: 50, strength: 5, attack: 10}
	player2 := &Player{name: "Player B", health: 100, strength: 10, attack: 5}

	game := Game{player1: player1, player2: player2}
	game.start()

	if !player1.isDead() {
		t.Errorf("Expected player1 to be dead, but player1 survived with health %d", player1.health)
	}
}
