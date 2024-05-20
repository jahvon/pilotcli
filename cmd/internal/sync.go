package internal

import (
	"github.com/spf13/cobra"

	"github.com/jahvon/flow/cmd/internal/interactive"
	"github.com/jahvon/flow/config/cache"
	"github.com/jahvon/flow/internal/context"
)

func RegisterSyncCmd(ctx *context.Context, rootCmd *cobra.Command) {
	subCmd := &cobra.Command{
		Use:   "sync",
		Short: "Scan workspaces and update flow cache.",
		Args:  cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			interactive.InitInteractiveCommand(ctx, cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			syncFunc(ctx, cmd, args)
		},
	}
	rootCmd.AddCommand(subCmd)
}

func syncFunc(ctx *context.Context, _ *cobra.Command, _ []string) {
	logger := ctx.Logger
	if err := cache.UpdateAll(ctx.Logger); err != nil {
		logger.FatalErr(err)
	}
	logger.PlainTextSuccess("Synced flow cache")
}
