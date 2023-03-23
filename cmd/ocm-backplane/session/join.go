package session

import (
	"fmt"

	"github.com/spf13/cobra"
)

var JoinCmd = &cobra.Command{
	Use: "join [SESSION_NAME]",
	Short: "Join an existing backplane session",
	Long: "Join an existing backplane session by initializing a new $SHELL instance in an isolated directory an pre-populating it's environment with useful variables",
	Args: cobra.RangeArgs(0, 1),
	RunE: runJoin,
	DisableAutoGenTag: true,
}

func runJoin(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		// This can be determined by 
		// 	- mtime of the session resources (what, specifically, would depend on the implementation, but most likely whatever dir encapsulates the session itself)
		//      - a dotfile in the session space itself, containing the last accessed time
		// 	- store the session info in ~/.config/backplane, ~/.cache/backplane, etc
		fmt.Println("Joining the last created session.")
		return nil
	}
	fmt.Println("joining a session named ", args[0])
	return nil
}
