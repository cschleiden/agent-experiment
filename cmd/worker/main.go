package main

import (
	"context"
	"fmt"

	"github.com/cschleiden/agent-experiment"
	"github.com/cschleiden/go-workflows/backend/sqlite"
	"github.com/cschleiden/go-workflows/worker"
)

func realMain() error {
	b := sqlite.NewInMemoryBackend()
	worker := worker.New(b, &worker.DefaultOptions)

	// client := client.New(b)

	worker.RegisterWorkflow(agent.SimpleAgent)

	ctx := context.Background()
	if err := worker.Start(ctx); err != nil {
		return fmt.Errorf("starting worker: %w", err)
	}

	// Wait for completion

	return nil
}

func main() {
	if err := realMain(); err != nil {
		panic(err)
	}
}
