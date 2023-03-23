package session

import "github.com/spf13/cobra"

var SessionCmd = &cobra.Command{
	Use: "session",
	Short: "Initialize, delete, join and manage isolated cluster sessions",
	Args: cobra.NoArgs,
	DisableAutoGenTag: true,
	Run: help,
}

func init() {
	SessionCmd.AddCommand(CreateCmd)
	SessionCmd.AddCommand(JoinCmd)
	SessionCmd.AddCommand(DeleteCmd)
}

func help(cmd *cobra.Command, _ []string) {
	_ = cmd.Help()
}
