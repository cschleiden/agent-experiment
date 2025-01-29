package main

import (
	"context"
	"fmt"

	"github.com/cschleiden/agent-experiment"
	"github.com/cschleiden/go-workflows/backend/sqlite"
	"github.com/cschleiden/go-workflows/client"
)

func realMain() error {
	// TODO: Extract backend creation
	b := sqlite.NewInMemoryBackend()

	c := client.New(b)
	ctx := context.Background()

	// TODO: Parse some input arguments for our new agent

	// TODO: Pass arguments to agent workflow
	_, err := c.CreateWorkflowInstance(ctx, client.WorkflowInstanceOptions{}, agent.SimpleAgent)

	if err != nil {
		return fmt.Errorf("creating workflow instance: %w", err)
	}

	return nil
}

func main() {
	if err := realMain(); err != nil {
		panic(err)
	}
}
