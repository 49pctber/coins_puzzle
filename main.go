package main

import (
	"log"
	"math/rand"
)

const N_COINS int = 100
const N_SIMS int = 1000000

func main() {
	log.Println("Two heads")
	runSimulation(N_SIMS, runGame)

	log.Println("Two-in-a-row")
	runSimulation(N_SIMS, runGame2)

	log.Println("Two-in-a-row heads *or* tails")
	runSimulation(N_SIMS, runGame3)
}

func runSimulation(n int, game func() (bool, bool)) {

	alice_wins := 0
	bob_wins := 0
	ties := 0

	for range n {
		alice_win, bob_win := game()

		if alice_win && bob_win {
			ties++
		} else if alice_win {
			alice_wins++
		} else {
			bob_wins++
		}
	}

	log.Printf("Alice wins: %d\n", alice_wins)
	log.Printf("Bob wins  : %d\n", bob_wins)
	log.Printf("Ties      : %d\n", ties)
}

/*
Game where the first player to find two heads (not necessarily consecutively) wins
*/
func runGame() (bool, bool) {

	// flip coins
	coins := make([]bool, N_COINS)
	for i := range coins {
		coins[i] = rand.Int()%2 == 0
	}

	// apply strategies
	alice_count := 0
	bob_count := 0

	for i := range N_COINS {
		if coins[i] {
			alice_count++
		}

		if coins[2*i+i/100] {
			bob_count++
		}

		if alice_count == 2 || bob_count == 2 {
			break
		}
	}

	return alice_count == 2, bob_count == 2
}

/*
Modified "two-heads-in-a-row" version where players must find two heads in a row to win.
*/
func runGame2() (bool, bool) {

	// flip coins
	coins := make([]bool, N_COINS)
	for i := range coins {
		coins[i] = rand.Int()%2 == 0
	}

	// apply strategies
	alice_count := 0
	bob_count := 0

	for i := range N_COINS {
		if coins[i] {
			alice_count++
		} else {
			alice_count = 0
		}

		if coins[2*i+i/100] {
			bob_count++
		} else {
			bob_count = 0
		}

		if alice_count == 2 || bob_count == 2 {
			break
		}
	}

	return alice_count == 2, bob_count == 2
}

/*
Modified "two-in-a-row" version where players need to see either two heads or two tails in a row to win.
*/
func runGame3() (bool, bool) {

	// flip coins
	coins := make([]bool, N_COINS)
	for i := range coins {
		coins[i] = rand.Int()%2 == 0
	}

	// apply strategies
	alice_count_heads := 0
	alice_count_tails := 0
	bob_count_heads := 0
	bob_count_tails := 0

	for i := range N_COINS {
		if coins[i] {
			alice_count_heads++
			alice_count_tails = 0
		} else {
			alice_count_heads = 0
			alice_count_tails++
		}

		if coins[2*i+i/100] {
			bob_count_heads++
			bob_count_tails = 0
		} else {
			bob_count_heads = 0
			bob_count_tails++
		}

		if alice_count_heads == 2 || alice_count_tails == 2 || bob_count_heads == 2 || bob_count_tails == 2 {
			break
		}
	}

	return alice_count_heads == 2 || alice_count_tails == 2, bob_count_heads == 2 || bob_count_tails == 2
}
