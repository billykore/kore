package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/billykore/kore/kore-cli/tpl"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const createDesc = `
This command create new service inside the monorepo services directory.

For example, 'kore-cli create todo' will create a service inside the 'services/' directory
that look something like this:

    kore/
    ├── ...
    ├── services/
    │   └── todo/
    │       ├── cmd/         # Contains main.go and wire.go injector files.
    │       ├── deployment/  # Kubernetes deployment configs.
    │       ├── repository/  # Service repositories.
    │       ├── server/      # Service http and gRPC servers.
    │       ├── service/     # Service API handlers.
    │       ├── usecase/     # Service usecases. Contains business logics.
    │       └── Dockerfile
    ├── ...

'kore-cli create' take service name for an argument and the name will be same for new directory
inside the 'services/' directory.
`

func newCreateCmd() *cobra.Command {
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
			structName := cases.Title(language.English).String(svcName)

			o := &createOption{
				AbsolutePath: wd,
				ServiceName:  svcName,
				StructName:   structName,
				Mod:          mod,
			}

			return o.Create()
		},
	}

	return cmd
}

type createOption struct {
	AbsolutePath string
	ServiceName  string
	StructName   string
	Mod          string
}

func (s *createOption) Create() error {
	svcPath := s.AbsolutePath + "/services"

	// check if services dir exist
	if _, err := os.Stat(svcPath); os.IsNotExist(err) {
		return err
	}
	svcPath = fmt.Sprintf("%s/%s", svcPath, s.ServiceName)

	// create service folder
	if err := os.Mkdir(svcPath, 0754); err != nil {
		return err
	}

	if err := s.createRepository(svcPath); err != nil {
		return err
	}
	if err := s.createUsecase(svcPath); err != nil {
		return err
	}
	if err := s.createService(svcPath); err != nil {
		return err
	}
	if err := s.createServer(svcPath); err != nil {
		return err
	}
	if err := s.createCmd(svcPath); err != nil {
		return err
	}
	if err := s.createDockerfile(svcPath); err != nil {
		return err
	}

	return nil
}

func (s *createOption) createRepository(path string) error {
	repoPath := fmt.Sprintf("%s/repository", path)

	if err := os.Mkdir(repoPath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", repoPath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.RepoProviderTemplate())))
	if err := providerTpl.Execute(providerFile, s); err != nil {
		return err
	}

	repoFile, err := os.Create(fmt.Sprintf("%s/greet_repository.go", repoPath))
	if err != nil {
		return err
	}
	repoTpl := template.Must(template.New("greet_repository").Parse(string(tpl.RepoTemplate())))
	if err := repoTpl.Execute(repoFile, s); err != nil {
		return err
	}

	return nil
}

func (s *createOption) createUsecase(path string) error {
	usecasePath := fmt.Sprintf("%s/usecase", path)

	if err := os.Mkdir(usecasePath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", usecasePath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.UsecaseProviderTemplate())))
	if err := providerTpl.Execute(providerFile, s); err != nil {
		return err
	}

	usecaseFile, err := os.Create(fmt.Sprintf("%s/greet_usecase.go", usecasePath))
	if err != nil {
		return err
	}
	usecaseTpl := template.Must(template.New("greet_usecase").Parse(string(tpl.UsecaseTemplate())))
	if err := usecaseTpl.Execute(usecaseFile, s); err != nil {
		return err
	}

	return nil
}

func (s *createOption) createService(path string) error {
	servicePath := fmt.Sprintf("%s/service", path)

	if err := os.Mkdir(servicePath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", servicePath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.ServiceProviderTemplate())))
	if err := providerTpl.Execute(providerFile, s); err != nil {
		return err
	}

	serviceFile, err := os.Create(fmt.Sprintf("%s/greet_service.go", servicePath))
	if err != nil {
		return err
	}
	serviceTpl := template.Must(template.New("greet_service").Parse(string(tpl.ServiceTemplate())))
	if err := serviceTpl.Execute(serviceFile, s); err != nil {
		return err
	}

	return nil
}

func (s *createOption) createServer(path string) error {
	servicePath := fmt.Sprintf("%s/server", path)

	if err := os.Mkdir(servicePath, 0754); err != nil {
		return err
	}

	providerFile, err := os.Create(fmt.Sprintf("%s/provider.go", servicePath))
	if err != nil {
		return err
	}
	providerTpl := template.Must(template.New("provider").Parse(string(tpl.ServerProviderTemplate())))
	if err := providerTpl.Execute(providerFile, s); err != nil {
		return err
	}

	httpFile, err := os.Create(fmt.Sprintf("%s/http.go", servicePath))
	if err != nil {
		return err
	}
	httpTpl := template.Must(template.New("http").Parse(string(tpl.HTTPServerTemplate())))
	if err := httpTpl.Execute(httpFile, s); err != nil {
		return err
	}

	grpcFile, err := os.Create(fmt.Sprintf("%s/grpc.go", servicePath))
	if err != nil {
		return err
	}
	grpcTpl := template.Must(template.New("grpc").Parse(string(tpl.GRPCServerTemplate())))
	if err := grpcTpl.Execute(grpcFile, s); err != nil {
		return err
	}

	return nil
}

func (s *createOption) createCmd(path string) error {
	servicePath := fmt.Sprintf("%s/cmd", path)

	if err := os.Mkdir(servicePath, 0754); err != nil {
		return err
	}

	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", servicePath))
	if err != nil {
		return err
	}
	mainTpl := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	if err := mainTpl.Execute(mainFile, s); err != nil {
		return err
	}

	wireFile, err := os.Create(fmt.Sprintf("%s/wire.go", servicePath))
	if err != nil {
		return err
	}
	wireTpl := template.Must(template.New("wire").Parse(string(tpl.WireTemplate())))
	if err := wireTpl.Execute(wireFile, s); err != nil {
		return err
	}

	return nil
}

func (s *createOption) createDockerfile(path string) error {
	dockerfile, err := os.Create(fmt.Sprintf("%s/Dockerfile", path))
	if err != nil {
		return err
	}
	dockerfileTpl := template.Must(template.New("Dockerfile").Parse(string(tpl.DockerfileTemplate())))
	if err := dockerfileTpl.Execute(dockerfile, s); err != nil {
		return err
	}
	return nil
}

func getMod() (string, error) {
	goMod, err := exec.Command("go", "list", "-m").Output()
	return strings.TrimSpace(string(goMod)), err
}
