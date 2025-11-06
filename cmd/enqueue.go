package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	// We will create these packages next
	//"github.com/PAAR16/queuectl/internal/job"
	//"github.com/PAAR16/queuectl/internal/store"
)

func init() {
	rootCmd.AddCommand(enqueueCmd)
}

var enqueueCmd = &cobra.Command{
	Use:   "enqueue [json-spec]",
	Short: "Enqueue a new job",
	Long:  `Adds a new job to the queue. The job must be specified in a JSON string.`,
	Args:  cobra.ExactArgs(1), // We expect exactly one argument
	RunE: func(cmd *cobra.Command, args []string) error {
		var newJob job.Job
		jobSpec := args[0]

		// Unmarshal the user-provided JSON into our job struct
		if err := json.Unmarshal([]byte(jobSpec), &newJob); err != nil {
			return fmt.Errorf("invalid job JSON: %w", err)
		}

		// --- Set Defaults ---
		if newJob.ID == "" {
			newJob.ID = uuid.New().String()
		}
		newJob.State = job.StatePending
		newJob.Attempts = 0
		if newJob.MaxRetries == 0 {
			newJob.MaxRetries = 3 // Default value
		}
		now := time.Now().UTC()
		newJob.CreatedAt = now
		newJob.UpdatedAt = now

		// --- Persist the Job ---
		jobStore, err := store.NewFileStore()
		if err != nil {
			return fmt.Errorf("failed to initialize job store: %w", err)
		}

		if err := jobStore.AddJob(newJob); err != nil {
			return fmt.Errorf("failed to add job: %w", err)
		}

		fmt.Printf("Successfully enqueued job with ID: %s\n", newJob.ID)
		return nil
	},
}