package executor

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/thesouldev/goboxd/internal/models"
)

const executionTimeout = 5 * time.Second

func Run(
	filePath string,
	stdin string,
) (*models.RunResponse, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		executionTimeout,
	)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"python3",
		filePath,
	)

	cmd.Stdin = strings.NewReader(stdin)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// Timeout
	if ctx.Err() == context.DeadlineExceeded {
		return &models.RunResponse{
			Status: "time_exceeded",
			Error:  "time_exceeded",
		}, nil
	}

	// Runtime Error
	if err != nil {
		return &models.RunResponse{
			Status: "runtime_error",
			Stdout: stdout.String(),
			Stderr: stderr.String(),
			Error:  "runtime_error",
		}, nil
	}

	// Success
	return &models.RunResponse{
		Status: "accepted",
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}, nil
}