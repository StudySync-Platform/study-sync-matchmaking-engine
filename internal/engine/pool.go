package engine

import (
	"errors"
	"sync"

	"studysync-matchmaking-engine/internal/matchmakingpb"
)

// MatchCandidate wraps the request along with internal metadata if needed
type MatchCandidate struct {
	Request *matchmakingpb.MatchRequest
}

// MatchPool stores candidates waiting for a match
type MatchPool struct {
	mu         sync.Mutex
	candidates []*MatchCandidate
}

// NewMatchPool initializes an empty pool
func NewMatchPool() *MatchPool {
	return &MatchPool{
		candidates: []*MatchCandidate{},
	}
}

// Add inserts a new candidate into the matchmaking pool
func (p *MatchPool) Add(req *matchmakingpb.MatchRequest) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.candidates = append(p.candidates, &MatchCandidate{Request: req})
	return nil
}

// FindBestMatch searches for a compatible match
func (p *MatchPool) FindBestMatch(req *matchmakingpb.MatchRequest) (*matchmakingpb.MatchRequest, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, candidate := range p.candidates {
		// Example simple match logic: same subject and level
		if candidate.Request.Subject == req.Subject &&
			candidate.Request.Level == req.Level &&
			candidate.Request.UserId != req.UserId {

			// Remove matched candidate from pool
			p.candidates = append(p.candidates[:i], p.candidates[i+1:]...)

			// Return matched request
			return candidate.Request, nil
		}
	}

	// No match found
	return nil, errors.New("no match found yet")
}
