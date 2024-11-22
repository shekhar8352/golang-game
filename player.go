package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name     string
	health   int
	strength int
	attack   int
}


func rollDie() int {
	return rand.Intn(6) + 1
}

func (p *Player) attackPlayer(opponent *Player) {
	// attackPlayer performs an attack on the given opponent Player.
	// It calculates the attack and defense rolls using rollDie, computes the damage dealt,
	// and reduces the opponent's health accordingly. If either player is nil, the function
	// prints an error message and exits. It also prints the details of the attack and the 
	// resulting health of the opponent.
	if p == nil {
		fmt.Println("p is nil")
		return
	}
	if opponent == nil {
		fmt.Println("opponent is nil")
		return
	}

	attackRoll := rollDie()
	defenseRoll := rollDie()

	attackDamage := p.attack * attackRoll
	defense := opponent.strength * defenseRoll

	totalDamage := attackDamage - defense
	if totalDamage > 0 {
		opponent.health -= totalDamage
	} else {
		totalDamage = 0
	}

	fmt.Printf("%s attacks %s! Attack roll: %d, Defense roll: %d\n", p.name, opponent.name, attackRoll, defenseRoll)
	fmt.Printf("Damage dealt: %d, %s's health reduced to: %d\n", totalDamage, opponent.name, opponent.health)
}

func (p *Player) isDead() bool {
	// Returns true if the player is dead.
	return p.health <= 0
}