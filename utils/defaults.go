package utils

import (
	"github.com/spf13/cobra"
)

func GetPStringFlag(rootCmd *cobra.Command, flagName string) (string, error) {
	flag, err := rootCmd.PersistentFlags().GetString(flagName)
	if err != nil {
		return "", err
	}
	return flag, nil
}
