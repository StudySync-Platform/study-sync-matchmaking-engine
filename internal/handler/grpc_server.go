package handler

import (
	"context"
	"log"

	"studysync-matchmaking-engine/internal/engine"
	"studysync-matchmaking-engine/internal/matchmakingpb"
)

// MatchmakingServer implements the gRPC server logic
type MatchmakingServer struct {
	matchmakingpb.UnimplementedMatchmakingServiceServer
	MatchEngine *engine.MatchEngine
}

// NewGRPCServer creates a new gRPC server with the matchmaking engine injected
func NewGRPCServer(engine *engine.MatchEngine) *MatchmakingServer {
	return &MatchmakingServer{
		MatchEngine: engine,
	}
}

// RequestMatch handles incoming matchmaking requests from clients
func (s *MatchmakingServer) RequestMatch(ctx context.Context, req *matchmakingpb.MatchRequest) (*matchmakingpb.MatchResponse, error) {
	// Try to find a match immediately
	match, err := s.MatchEngine.FindBestMatch(req)
	if err == nil && match != nil {
		log.Printf("✅ Match found: %d <-> %d", req.UserId, match.UserId)

		return &matchmakingpb.MatchResponse{
			Matched:       true,
			SessionId:     generateSessionID(req.UserId, match.UserId),
			Message:       "✅ Match found successfully",
			MatchedUserId: match.UserId,
		}, nil
	}

	// If no match, add the user to the pool
	_ = s.MatchEngine.AddToPool(req)

	return &matchmakingpb.MatchResponse{
		Matched: false,
		Message: "🕓 Waiting for a match...",
	}, nil
}

// generateSessionID creates a deterministic session ID between two users
func generateSessionID(user1, user2 int64) int64 {
	if user1 < user2 {
		return user1*100000 + user2
	}
	return user2*100000 + user1
}
