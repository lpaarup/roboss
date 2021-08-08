package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "roboss",
		Short: "Roboss is a robotics orchestration platform",
		Long: `With roboss you can manage workers and topics on top of
		kubernetes.`,
	}
)

func init() {
	rootCmd.AddCommand(coreCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
