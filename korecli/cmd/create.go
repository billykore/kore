package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/billykore/kore/korecli/tpl"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const createDesc = `
This command create new service inside the monorepo services directory.

Use --repo or -r flag for create new repository interface in the libs/repository directory.
E.g.  korecli create todo --repo Todo

For example, 'korecli create todo' will create a service inside the services/ directory 
and a protobuf inside the  file that look something like this:

    kore/
    ├──...
    ├── libs/
    │   ├──...
    │   ├── proto/
    │   │   └── v1/
    │   │       ├──...
    │   │       ├── todo.proto
    │   │       └──...
    │   ├──...
    │   └── repository
    │       ├──...
    │       ├── todo.go
    │       └──...
    ├── services/
    │   └── todo/
    │       ├── cmd/         # Contains main.go and wire.go injector files.
    │       ├── deployment/  # Kubernetes deployment configs.
    │       ├── repository/  # Service repositories.
    │       ├── server/      # Service http and gRPC servers.
    │       ├── service/     # Service API handlers.
    │       ├── usecase/     # Service usecases.
    │       └── Dockerfile
    └──...

'korecli create' take service name for an argument and the name will be same for new directory
inside the services/ directory.
`

func newCreateCmd() *cobra.Command {
	d := &createData{}

	cmd := &cobra.Command{
		Use:   "create SERVICE",
		Short: "Create new service",
		Long:  createDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			mod, err := getMod()
			if err != nil {
				return err
			}

			svcName := args[0]
			d.AbsolutePath = wd
			d.ServiceName = svcName
			d.StructName = cases.Title(language.English).String(svcName)
			d.Mod = mod

			return d.create()
		},
	}

	f := cmd.Flags()
	f.StringVarP(&d.Repository, "repo", "r", "", "create new repository")

	return cmd
}

type createData struct {
	AbsolutePath string
	ServiceName  string
	StructName   string
	Mod          string
	Repository   string
}

func (d *createData) create() error {
	libsPath := d.AbsolutePath + "/libs"
	// check if libs dir exist
	if _, err := os.Stat(libsPath); os.IsNotExist(err) {
		return err
	}
	// create proto
	if err := d.createProto(libsPath); err != nil {
		return err
	}
	// create repository interface (if required)
	if err := d.createIRepository(libsPath); err != nil {
		return err
	}

	svcPath := d.AbsolutePath + "/services"
	// check if services dir exist
	if _, err := os.Stat(svcPath); os.IsNotExist(err) {
		return err
	}
	svcPath = fmt.Sprintf("%s/%s", svcPath, d.ServiceName)

	// create service
	if err := os.Mkdir(svcPath, 0754); err != nil {
		return err
	}
	if err := d.createRepository(svcPath); err != nil {
		return err
	}
	if err := d.createUsecase(svcPath); err != nil {
		return err
	}
	if err := d.createService(svcPath); err != nil {
		return err
	}
	if err := d.createServer(svcPath); err != nil {
		return err
	}
	if err := d.createCmd(svcPath); err != nil {
		return err
	}
	if err := d.createDeployment(svcPath); err != nil {
		return err
	}
	if err := d.createDockerfile(svcPath); err != nil {
		return err
	}

	return nil
}

func (d *createData) createProto(path string) error {
	protoPath := fmt.Sprintf("%s/proto/v1", path)

	protoFile, err := os.Create(fmt.Sprintf("%s/%s.proto", protoPath, d.ServiceName))
	if err != nil {
		return err
	}
	protoTpl := template.Must(template.New(d.ServiceName).Parse(string(tpl.ProtoTemplate())))
	if err := protoTpl.Execute(protoFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createRepository(path string) error {
	repoPath := fmt.Sprintf("%s/repository", path)

	if err := os.Mkdir(repoPath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", repoPath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.RepoProviderTemplate())))
	if err := providerTpl.Execute(providerFile, d); err != nil {
		return err
	}

	repoFile, err := os.Create(fmt.Sprintf("%s/%s_repository.go", repoPath, d.ServiceName))
	if err != nil {
		return err
	}
	repoTpl := template.Must(template.New("repository").Parse(string(tpl.RepoTemplate())))
	if err := repoTpl.Execute(repoFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createUsecase(path string) error {
	usecasePath := fmt.Sprintf("%s/usecase", path)

	if err := os.Mkdir(usecasePath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", usecasePath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.UsecaseProviderTemplate())))
	if err := providerTpl.Execute(providerFile, d); err != nil {
		return err
	}

	usecaseFile, err := os.Create(fmt.Sprintf("%s/%s_usecase.go", usecasePath, d.ServiceName))
	if err != nil {
		return err
	}
	usecaseTpl := template.Must(template.New("usecase").Parse(string(tpl.UsecaseTemplate())))
	if err := usecaseTpl.Execute(usecaseFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createService(path string) error {
	servicePath := fmt.Sprintf("%s/service", path)

	if err := os.Mkdir(servicePath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", servicePath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.ServiceProviderTemplate())))
	if err := providerTpl.Execute(providerFile, d); err != nil {
		return err
	}

	serviceFile, err := os.Create(fmt.Sprintf("%s/%s_service.go", servicePath, d.ServiceName))
	if err != nil {
		return err
	}
	serviceTpl := template.Must(template.New("service").Parse(string(tpl.ServiceTemplate())))
	if err := serviceTpl.Execute(serviceFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createServer(path string) error {
	serverPath := fmt.Sprintf("%s/server", path)

	if err := os.Mkdir(serverPath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", serverPath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.ServerProviderTemplate())))
	if err := providerTpl.Execute(providerFile, d); err != nil {
		return err
	}

	httpFile, err := os.Create(fmt.Sprintf("%s/http.go", serverPath))
	if err != nil {
		return err
	}
	httpTpl := template.Must(template.New("http").Parse(string(tpl.HTTPServerTemplate())))
	if err := httpTpl.Execute(httpFile, d); err != nil {
		return err
	}

	grpcFile, err := os.Create(fmt.Sprintf("%s/grpc.go", serverPath))
	if err != nil {
		return err
	}
	grpcTpl := template.Must(template.New("grpc").Parse(string(tpl.GRPCServerTemplate())))
	if err := grpcTpl.Execute(grpcFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createCmd(path string) error {
	cmdPath := fmt.Sprintf("%s/cmd", path)

	if err := os.Mkdir(cmdPath, 0754); err != nil {
		return err
	}

	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", cmdPath))
	if err != nil {
		return err
	}
	mainTpl := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	if err := mainTpl.Execute(mainFile, d); err != nil {
		return err
	}

	wireFile, err := os.Create(fmt.Sprintf("%s/wire.go", cmdPath))
	if err != nil {
		return err
	}
	wireTpl := template.Must(template.New("wire").Parse(string(tpl.WireTemplate())))
	if err := wireTpl.Execute(wireFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createDeployment(path string) error {
	deploymentPath := fmt.Sprintf("%s/deployment", path)

	if err := os.Mkdir(deploymentPath, 0754); err != nil {
		return err
	}

	deploymentFile, err := os.Create(fmt.Sprintf("%s/deployment.yaml", deploymentPath))
	if err != nil {
		return err
	}
	deploymentTpl := template.Must(template.New("main").Parse(string(tpl.DeploymentTemplate())))
	if err := deploymentTpl.Execute(deploymentFile, d); err != nil {
		return err
	}

	envFile, err := os.Create(fmt.Sprintf("%s/env.yaml", deploymentPath))
	if err != nil {
		return err
	}
	envTpl := template.Must(template.New("main").Parse(string(tpl.EnvTemplate())))
	if err := envTpl.Execute(envFile, d); err != nil {
		return err
	}

	return nil
}

func (d *createData) createDockerfile(path string) error {
	dockerfile, err := os.Create(fmt.Sprintf("%s/Dockerfile", path))
	if err != nil {
		return err
	}
	dockerfileTpl := template.Must(template.New("Dockerfile").Parse(string(tpl.DockerfileTemplate())))
	if err := dockerfileTpl.Execute(dockerfile, d); err != nil {
		return err
	}
	return nil
}

func (d *createData) createIRepository(path string) error {
	if d.Repository == "" {
		return nil
	}

	repoPath := path + "/repository"

	repoFile, err := os.Create(fmt.Sprintf("%s/%s.go", repoPath, strings.ToLower(d.Repository)))
	if err != nil {
		return err
	}
	repoTpl := template.Must(template.New("repository").Parse(string(tpl.IRepositoryTemplate())))
	if err := repoTpl.Execute(repoFile, d); err != nil {
		return err
	}

	return nil
}

func getMod() (string, error) {
	goMod, err := exec.Command("go", "list", "-m").Output()
	return strings.TrimSpace(string(goMod)), err
}
