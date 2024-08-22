package cmd

import (
	"github.com/anosora233/ghstats/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(repoCmd)
}

var repoCmd = &cobra.Command{
	Use:   "repo <user>",
	Args:  cobra.ExactArgs(1),
	Short: "Show repositories of the user",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.DisplayRepositories(args[0])
	},
}
