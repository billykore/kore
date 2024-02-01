package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const koreDesc = `Monorepo CLI tool

Action for CLI:

- korecli create: create new service
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

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
