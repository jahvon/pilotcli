// Package cmd handle the cli commands
package cmd

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/jahvon/tbox/internal/cmd/version"
	"github.com/jahvon/tbox/internal/io"
)

var rootCmd = &cobra.Command{
	Use:   "tbox",
	Short: "[Alpha] CLI script wrapper",
	Long:  `Command line interface script wrapper`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verbose, err := strconv.ParseBool(cmd.Flag("verbose").Value.String())
		if err != nil {
			io.PrintErrorAndExit(fmt.Errorf("invalid verbose flag - %v", err))
		}

		if verbose {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
			io.PrintInfo("Verbose logging enabled")
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	},
	Version: version.String(),
}

var log = io.Log()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		io.PrintErrorAndExit(fmt.Errorf("failed to execute command: %w", err))
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose log output")
}
