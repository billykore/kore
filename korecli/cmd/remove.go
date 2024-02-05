package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const removeDesc = `
This command remove a service from the monorepo 'services/'.
Also will delete the protobuf and all proto generated files.

For example, 'korecli rm todo' will remove 'service/todo' directory
and '/libs/proto/v1/todo.proto' file with all the proto generated files.
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
	if err := d.removeProto(); err != nil {
		return err
	}
	if err := d.removeService(); err != nil {
		return err
	}
	return nil
}

func (d *removeData) removeProto() error {
	protoPath := d.AbsolutePath + "/libs/proto/v1"
	if _, err := os.Stat(protoPath); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s.proto", protoPath, d.ServiceName)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s.pb.go", protoPath, d.ServiceName)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s.pb.gw.go", protoPath, d.ServiceName)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s_grpc.pb.go", protoPath, d.ServiceName)); err != nil {
		return err
	}
	return nil
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
