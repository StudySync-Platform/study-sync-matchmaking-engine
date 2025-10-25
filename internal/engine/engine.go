package engine

import (
	"studysync-matchmaking-engine/internal/matchmakingpb"
)

// MatchEngine is responsible for finding matches and managing the pool
type MatchEngine struct {
	pool *MatchPool
}

// NewMatchEngine creates a new engine instance with a given pool
func NewMatchEngine(pool *MatchPool) *MatchEngine {
	return &MatchEngine{
		pool: pool,
	}
}

// AddToPool adds a new candidate to the pool for future matching
func (e *MatchEngine) AddToPool(req *matchmakingpb.MatchRequest) error {
	return e.pool.Add(req)
}

// FindBestMatch attempts to find a compatible match for a given request
func (e *MatchEngine) FindBestMatch(req *matchmakingpb.MatchRequest) (*matchmakingpb.MatchRequest, error) {
	return e.pool.FindBestMatch(req)
}
