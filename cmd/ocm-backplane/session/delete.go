package session

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use: "delete [SESSION_NAME]",
	Short: "Remove a session",
	Long: "Remove a session and all corresponding objects",
	Args: cobra.RangeArgs(0, 1),
	RunE: runDelete,
	DisableAutoGenTag: true,
}

func init() {
	flags := DeleteCmd.Flags()
	flags.BoolVarP(
		&deleteAll,
		"all",
		"A",
		false,
		"Delete all sessions. Cannot be passed if specifying a SESSION_NAME",
	)
}

var (
	deleteAll bool
)

func runDelete(cmd *cobra.Command, args []string) error {
	if deleteAll {
		if len(args) != 0 {
			return fmt.Errorf("Cannot specify --all/-A with a session")
		}
		fmt.Println("Deleting all sessions")
		return nil
	}
	if len(args) == 0 {
		return fmt.Errorf("No session specified. Nothing to do.")
	}

	fmt.Printf("Deleting session %s\n", args[0])
	return nil
}
