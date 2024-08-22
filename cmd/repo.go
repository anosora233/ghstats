package cmd

import (
	"github.com/anosora233/ghstats/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(repoCmd)
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Show repositories of the user",
	PreRun: func(cmd *cobra.Command, args []string) {
		rootCmd.MarkPersistentFlagRequired("username")
	},
	Run: func(cmd *cobra.Command, args []string) {
		pkg.DisplayRepositories(username)
	},
}
