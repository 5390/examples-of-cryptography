package hyperledger

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Validator represents a Solana validator node
// Each validator has:
// - ID: A unique identifier for the validator.
// - Vote: A simulated vote (0 or 1) indicating "no" or "yes".
type Validator struct {
	ID   int // Validator ID
	Vote int // Vote result (0 or 1)
}

// SimulateLeaderRotation simulates the process of rotating leaders among validators.
// Each validator takes turns being the leader, generating a block (Proof of History simulation).
// Parameters:
// - validators: A slice of Validator structs representing the participating nodes.
func SimulateLeaderRotation(validators []Validator) {
	var wg sync.WaitGroup // Used to manage concurrency and wait for all goroutines to finish.

	// Iterate through each validator to assign them as a leader in rotation.
	for i := 0; i < len(validators); i++ {
		wg.Add(1) // Increment the WaitGroup counter before starting the goroutine.

		// Launch a goroutine for each leader to simulate block generation.
		go func(leader Validator) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes.
			fmt.Printf("Validator %d is the leader. Generating PoH...\n", leader.ID)
			// Simulate block processing time using a random delay.
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Validator %d completed block proposal.\n", leader.ID)
		}(validators[i]) // Pass the validator as a parameter to the goroutine.

		time.Sleep(500 * time.Millisecond) // Simulate a leader rotation interval between leaders.
	}

	// Wait for all goroutines (leader tasks) to complete.
	wg.Wait()
}

// SimulateVoting simulates a voting process where validators cast votes (yes/no) on a block.
// Each validator generates a random vote value (0 or 1).
// Parameters:
// - validators: A slice of Validator structs representing the participating nodes.
func SimulateVoting(validators []Validator) {
	var wg sync.WaitGroup // Used to manage concurrency and wait for all goroutines to finish.

	// Iterate through each validator and simulate voting.
	for _, v := range validators {
		wg.Add(1) // Increment the WaitGroup counter before starting the goroutine.

		// Launch a goroutine for each validator to cast a vote.
		go func(validator Validator) {
			defer wg.Done()               // Decrement the WaitGroup counter when the goroutine finishes.
			validator.Vote = rand.Intn(2) // Simulate a random vote: 0 (no) or 1 (yes).
			fmt.Printf("Validator %d voted: %d\n", validator.ID, validator.Vote)
		}(v) // Pass the validator as a parameter to avoid closure issues.
	}

	// Wait for all goroutines (voting tasks) to complete.
	wg.Wait()
}

// SolanaExample is the entry point to simulate Solana's consensus process.
// It creates validators, rotates leaders, and simulates voting to reach consensus.
func SolanaExample() {
	// Create a set of validator nodes participating in the consensus process.
	validators := []Validator{
		{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, // Four validators with unique IDs.
	}

	// Begin the Solana consensus simulation.
	fmt.Println("Starting Solana Consensus Simulation...")

	// Simulate leader rotation, where validators take turns proposing blocks.
	SimulateLeaderRotation(validators)

	// Simulate voting on the proposed block.
	SimulateVoting(validators)

	// Print the conclusion of the consensus process.
	fmt.Println("Consensus process completed.")
}
