package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Short: "Generate and maintain configs",
	Args: cobra.NoArgs,
	DisableAutoGenTag: true,
	Run: help,
}

func init() {
	ConfigCmd.AddCommand(InitCmd)
	ConfigCmd.AddCommand(UpgradeCmd)
}

func help(cmd *cobra.Command, _ []string) {
	_ = cmd.Help()
}
