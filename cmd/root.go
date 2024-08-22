package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "ghstats"}

func Execute() {
	rootCmd.Execute()
}

var username string
var repository string

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "github username")
	rootCmd.PersistentFlags().StringVarP(&repository, "repository", "r", "", "github repository")
}
