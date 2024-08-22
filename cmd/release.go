package cmd

import (
	"github.com/anosora233/ghstats/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Show releases of the repository",
	PreRun: func(cmd *cobra.Command, args []string) {
		rootCmd.MarkPersistentFlagRequired("username")
		rootCmd.MarkPersistentFlagRequired("repository")
	},
	Run: func(cmd *cobra.Command, args []string) {
		pkg.DisplayReleases(username, repository)
	},
}
