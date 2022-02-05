package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of echoctl",
  Long:  `All software has versions. This is echoctl's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("echoctl micro echo server v0.1 -- HEAD")
  },
}