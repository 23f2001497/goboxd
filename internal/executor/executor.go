package executor

import (
	"bytes"
	"os/exec"
	"context"
	"time"

	"github.com/thesouldev/goboxd/internal/models"
)

const executionTimeout = 5*time.Second

func Run(filePath string) (*models.RunResponse, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		executionTimeout,
	)
	defer cancel()

	cmd := exec.Command("python3", filePath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	//Timeout check should come first
	if ctx.Err() == context.DeadlineExceeded {
		return &models.RunResponse{
			Error: "time_exceeded",
		}, nil
	}

	//Runtime error
	// if err !=nil{
	// 	return &models.RunResponse{
	// 		Stdout: stdout.String(),
	// 		Stderr: stderr.String(),
	// 		Error: fmt.Sprintf("runtime_error: %v", err),
	// 	}, nil
	// }
	if err != nil {
		return &models.RunResponse{
			Stdout: stdout.String(),
			Stderr: stderr.String(),
			Error:  "runtime_error",
		}, nil
	}
	//Success
	return &models.RunResponse{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}, nil
}
