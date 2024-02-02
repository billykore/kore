package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const koreDesc = `Monorepo CLI tool

Action for CLI:

- korecli create: create new service
- korecli rm:     remove a service
`

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "korecli",
		Short: "Monorepo CLI",
		Long:  koreDesc,
	}
	return cmd
}

func Execute() {
	root := newRootCmd()

	root.AddCommand(newCreateCmd())
	root.AddCommand(newRemoveCmd())

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
