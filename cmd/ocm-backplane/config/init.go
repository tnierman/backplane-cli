package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bootstrapRepo string
	bootstrapBranch string
)

var InitCmd = &cobra.Command{
	Use: "init",
	Short: "Initializes backplane config",
	Long: "Initializes backplane configuration files, optionally using a source a git repo as a source of truth to pre-populate the directory",
	Args: cobra.NoArgs,
	RunE: runInit,
	DisableAutoGenTag: true,
}

func init() {
	flags := InitCmd.Flags()
	flags.StringVarP(&bootstrapRepo, "repo", "r", "", "Specify a repo to bootstrap the configuration with")
	flags.StringVarP(&bootstrapBranch, "branch", "b", "master", "Specify a branch from the bootstrap repo to use. Must be invoked along with --repo/-r")
}

func runInit(cmd *cobra.Command, _ []string) error {
	if bootstrapRepo != "" {
		return bootstrapConfig()
	}

	return promptConfig()
}

func bootstrapConfig() error {
	fmt.Printf("I should pull down a repo '%s' a put branch '%s' into ~/.config/backplane\n", bootstrapRepo, bootstrapBranch)
	return nil
}

func promptConfig() error {
	fmt.Println("I should prompt the user to create basic config?")
	return nil
}
