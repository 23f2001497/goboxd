package executor

import (
	"bytes"
	"os/exec"

	"github.com/thesouldev/goboxd/internal/models"
)

func Run(filePath string) (*models.RunResponse, error) {
	cmd := exec.Command("python3", filePath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Run() // ignore exit error for now

	return &models.RunResponse{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}, nil
}
