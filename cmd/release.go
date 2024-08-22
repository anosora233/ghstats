package cmd

import (
	"github.com/anosora233/ghstats/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release <user> <repo>",
	Args:  cobra.ExactArgs(2),
	Short: "Show releases of the repository",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.DisplayReleases(args[0], args[1])
	},
}
