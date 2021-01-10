package cmd

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var longRunningWorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "long-running multiworker",
	RunE: func(cmd *cobra.Command, args []string) error {
		runLongRunningMultiworkerRunner()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(longRunningWorkerCmd)
}

func runLongRunningMultiworkerRunner() {
	forever := make(chan bool)

	for i := 0; i < 5; i++ {
		go worker(i)
	}

	<-forever
}

func worker(i int) {
	fmt.Printf("Started long-running job routine %v at %v\n", i, time.Now().UTC().Format(time.RFC3339Nano))
	for {
		fmt.Println(fmt.Sprintf("workerID: %v, timestamp: %v", i, time.Now().UTC().Format(time.RFC3339Nano)))
		time.Sleep(time.Second)
	}
}
