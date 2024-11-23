# Magical Arena Game

## Overview
**Magical Arena** is a turn-based combat simulation game where two players fight in an arena. Each player has `health`, `strength`, and `attack` attributes. Players attack and defend in turns using a dice roll mechanic, and the game ends when one player's health reaches 0.

## Table of Contents
- [Overview](#overview)
- [Gameplay](#gameplay)
- [Game Rules](#game-rules)
- [Setup](#setup)
- [Running the Game](#running-the-game)
- [Running Tests](#running-tests)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Future Improvements](#future-improvements)

## Gameplay
The game simulates combat between two players with the following attributes:
- **Health**: Represents the player's life. When it reaches 0, the player dies.
- **Strength**: Used in defense; it reduces the damage from incoming attacks.
- **Attack**: Represents the player's attacking power. Higher attack means more damage.

Players attack and defend by rolling dice:
- **Attacking Dice**: A 6-sided die rolled by the attacker.
- **Defending Dice**: A 6-sided die rolled by the defender.

### Example Combat
1. **Player A** rolls a 5 on the attacking die and has an attack of 10. Attack damage is calculated as `5 * 10 = 50`.
2. **Player B** rolls a 2 on the defending die and has a strength of 10. Defended damage is calculated as `2 * 10 = 20`.
3. Player B’s health is reduced by the difference (`50 - 20 = 30`).

Players continue attacking and defending in turns until one player's health reaches 0.

## Game Rules
1. Players take turns attacking and defending.
2. The player with the lower health attacks first.
3. The game ends when one player's health reaches 0.

## Setup
### Prerequisites
- Go (Golang) installed on your machine. [Download Go](https://golang.org/dl/)

### Installation
1. Clone the project:
   ```bash
   git clone <your-local-path>
   cd magical_arena

2. Install dependencies:
   ```bash
   go mod init

3. Build the project:
   ```bash
   go build

## Running the Game
1. Run the game with the following command:
   On windows:
   ```bash
   magical_arena.exe

   On Linux:
   ```bash
   ./magical_arena

## Running Tests
1. Run the tests with the following command:
   ```bash
   go test

## Project Structure
- **main.go**: The main entry point of the program.
- **player.go**: The player struct and its methods.
- **game.go**: # Contains Game struct and game logic.
- **project_test.go**: The test file for the project.

## Technologies Used
- **Go**: The programming language used for this project.
- **Go Modules**: The module system for Go.
- **Go Test**: The testing framework for Go.