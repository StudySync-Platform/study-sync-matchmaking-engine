package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"studysync-matchmaking-engine/internal/engine"
	"studysync-matchmaking-engine/internal/handler"
	"studysync-matchmaking-engine/internal/matchmakingpb"
)

func main() {
	// Start listening on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	// Initialize a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Create matchmaking pool and engine
	pool := engine.NewMatchPool()
	matchEngine := engine.NewMatchEngine(pool)

	// Register the gRPC server implementation with the engine
	matchmakingpb.RegisterMatchmakingServiceServer(
		grpcServer,
		handler.NewGRPCServer(matchEngine),
	)

	log.Println("✅ Matchmaking gRPC server running on port 50051")

	// Start serving
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
