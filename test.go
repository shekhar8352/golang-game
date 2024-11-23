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
	playerA := &Player{name: "Player A", health: 50, strength: 5, attack: 10}
	playerB := &Player{name: "Player B", health: 100, strength: 10, attack: 5}

	game := Game{player1: playerA, player2: playerB}
	game.start()

	if !playerA.isDead() && !playerB.isDead() {
		t.Errorf("One player should have died, but both are still alive.")
	}
}