package session

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use: "create [CLUSTERID|EXTERNAL_ID|CLUSTER_NAME|CLUSTER_NAME_SEARCH]",
	Short: "Create a new isolated cluster session",
	Long: "Initializes and joins a new isolated cluster session. Within the session, all commands automatically run against the cluster matching the given identifier.",
	Args: cobra.RangeArgs(0, 1),
	DisableAutoGenTag: true,
	RunE: runCreate,
}

var (
	// Flags
	sessionName string
	joinSession bool
)

func init() {
	flags := CreateCmd.Flags()
	flags.StringVarP(
		&sessionName,
		"name",
		"n",
		"",
		"Specify a session name. By default, sessions are named after the identifier provided; if no identifier provided, then the uuid of the cluster currently logged into is used",
	)
	flags.BoolVarP(
		&joinSession,
		"join",
		"j",
		true,
		"Join session automatically after creating. Defaults to true",
	)
}

func runCreate(cmd *cobra.Command, args []string) error {
	var clusterID string
	if len(args) == 0 {
		clusterID = "<current cluster>"
	} else {
		clusterID = args[0]
	}

	joinStatus := "joining"
	if !joinSession {
		joinStatus = "not joining"
	}

	if sessionName == "" {
		sessionName = clusterID
	}

	fmt.Printf("Creating isolated environment for the cluster '%s', naming the session '%s', and %s afterward\n", clusterID, sessionName, joinStatus)
	return nil
}
