package models

type RunRequest struct {
	Language string `json:"language"`
	Source   string `json:"source"`
	Stdin    string `json:"stdin,omitempty"`
}

type RunResponse struct {
	Status string `json:"status"`

	Stdout string `json:"stdout,omitempty"`
	Stderr string `json:"stderr,omitempty"`

	Error string `json:"error,omitempty"`
}