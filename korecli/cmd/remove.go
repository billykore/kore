package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const removeDesc = `
This command remove a service from the monorepo 'services/'.

For example, 'korecli rm todo' will remove 'service/todo' directory.
`

func newRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove SERVICE",
		Short:   "Remove a service",
		Long:    removeDesc,
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			d := &removeData{
				AbsolutePath: wd,
				ServiceName:  args[0],
			}

			return d.remove()
		},
	}
	return cmd
}

type removeData struct {
	AbsolutePath string
	ServiceName  string
}

func (d *removeData) remove() error {
	if err := d.removeEntity(); err != nil {
		return err
	}
	if err := d.removeService(); err != nil {
		return err
	}
	return nil
}

func (d *removeData) removeEntity() error {
	entityPath := d.AbsolutePath + "/pkg/entity"
	if _, err := os.Stat(entityPath); err != nil {
		return err
	}
	if err := removeIfExist(fmt.Sprintf("%s/%s.go", entityPath, d.ServiceName)); err != nil {
		return err
	}
	return nil
}

func removeIfExist(filename string) error {
	err := os.Remove(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return err
}

func (d *removeData) removeService() error {
	svcPath := fmt.Sprintf("%s/services/%s", d.AbsolutePath, d.ServiceName)
	if _, err := os.Stat(svcPath); err != nil {
		return err
	}
	if err := os.RemoveAll(svcPath); err != nil {
		return err
	}
	return nil
}
