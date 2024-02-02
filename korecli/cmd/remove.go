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
		Use:   "rm SERVICE",
		Short: "Remove a service",
		Long:  removeDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			o := &removeOption{
				AbsolutePath: wd,
				ServiceName:  args[0],
			}

			return o.remove()
		},
	}
	return cmd
}

type removeOption struct {
	AbsolutePath string
	ServiceName  string
}

func (o *removeOption) remove() error {
	if err := o.removeProto(); err != nil {
		return err
	}
	if err := o.removeService(); err != nil {
		return err
	}
	return nil
}

func (o *removeOption) removeProto() error {
	protoPath := o.AbsolutePath + "/libs/proto/v1"
	if _, err := os.Stat(protoPath); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s.proto", protoPath, o.ServiceName)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s.pb.go", protoPath, o.ServiceName)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s.pb.gw.go", protoPath, o.ServiceName)); err != nil {
		return err
	}
	if err := os.Remove(fmt.Sprintf("%s/%s_grpc.pb.go", protoPath, o.ServiceName)); err != nil {
		return err
	}
	return nil
}

func (o *removeOption) removeService() error {
	svcPath := fmt.Sprintf("%s/services/%s", o.AbsolutePath, o.ServiceName)
	if _, err := os.Stat(svcPath); err != nil {
		return err
	}
	if err := os.RemoveAll(svcPath); err != nil {
		return err
	}
	return nil
}
