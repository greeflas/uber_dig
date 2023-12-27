package app

import (
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	ctx := context.Background()

	if err := Run(ctx, EnvTest); err != nil {
		t.Fatal(err)
	}
}
