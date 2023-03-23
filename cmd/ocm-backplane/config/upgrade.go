package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpgradeCmd = &cobra.Command{
	Use: "upgrade",
	Short: "Upgrade the configuration",
	Long: "Upgrades the configuration. If the configuration source is 'git', this is accomplished by pulling down the latest config from the config dir's repo.",
	Args: cobra.NoArgs,
	RunE: runUpgrade,
	DisableAutoGenTag: true,
}

func runUpgrade(cmd *cobra.Command, args []string) error {
	fmt.Println("I should 'git pull' or 'git fetch && git rebase' ~/.config/backplane, if that directory has a .git/ dir. Otherwise, there's nothing to do")
	return nil
}
