package main

import (
	"fmt"
	"sync"
	"time"
)

type Participant struct {
	ID       string
	Decision bool
}

type Coordinator struct {
	Participants []*Participant
	mu           sync.Mutex
}

func NewParticipant(id string) *Participant {
	return &Participant{ID: id}
}

func NewCoordinator() *Coordinator {
	return &Coordinator{
		Participants: make([]*Participant, 0),
	}
}

func (p *Participant) Prepare() {
	// Simulate participant's preparation logic
	fmt.Printf("Participant %s: Prepared\n", p.ID)
	p.Decision = true // Simulate success
}

func (p *Participant) Commit() {
	// Simulate participant's commit logic
	fmt.Printf("Participant %s: Committed\n", p.ID)
}

func (p *Participant) Rollback() {
	// Simulate participant's rollback logic
	fmt.Printf("Participant %s: Rolled back\n", p.ID)
}

func (c *Coordinator) AddParticipant(participant *Participant) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Participants = append(c.Participants, participant)
}

func (c *Coordinator) PreparePhase() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Phase 1: Prepare
	fmt.Println("Phase 1: Prepare")
	for _, participant := range c.Participants {
		participant.Prepare()
	}

	// Simulate a decision based on the participants' success
	return true
}

func (c *Coordinator) CommitPhase(decision bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Phase 2: Commit or Rollback
	fmt.Println("Phase 2: Commit or Rollback")
	for _, participant := range c.Participants {
		if decision {
			participant.Commit()
		} else {
			participant.Rollback()
		}
	}
}

func main() {
	// Create participants
	participant1 := NewParticipant("Participant1")
	participant2 := NewParticipant("Participant2")

	// Create coordinator
	coordinator := NewCoordinator()

	// Add participants to the coordinator
	coordinator.AddParticipant(participant1)
	coordinator.AddParticipant(participant2)

	// Start the 2PC protocol
	decision := coordinator.PreparePhase()

	// Introduce a delay to simulate network communication and coordination
	time.Sleep(2 * time.Second)

	// Complete the 2PC protocol based on the decision
	coordinator.CommitPhase(decision)
}
