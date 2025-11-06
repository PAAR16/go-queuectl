package main

import (
	"time"
)

type state string

const (
	statepending state="pending"
	stateprocessing state="processing"
	statecompleted state="completed"
	statefailed state="failed"
	statedead state="dead"
)

type Job struct {
	ID           string    `json:"id"`
	Command      string    `json:"command"`
	State        state     `json:"state"`
	Attempts     int       `json:"attempts"`
	MaxRetries   int       `json:"max_retries"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	NextRunAt    time.Time `json:"next_run_at,omitempty"`
	ExecutionLog string    `json:"execution_log,omitempty"`
}



