package autorender

import (
	"path"
	"os"
	"syscall"
	"log"
)

type AutoRender struct {
	AerenderPath string
}

func New(aerenderPath string)(*AutoRender, error){
	return &AutoRender{
		AerenderPath: aerenderPath,
	},nil
}

func (render *AutoRender)Render(project *Project){

	project.Validate()
	project.Setup()
	project.DowloadAssets()
	project.ReplaceAssets()

	render.runAerender(project)

	project.ClearTempDir()
}

func (render *AutoRender)runAerender(project *Project){
	args := []string{render.AerenderPath,
		"-comp", project.Composition,
		"-project", path.Join(project.ProjectPath, project.ProjectName),
		"-output", path.Join(project.ProjectPath, "result"),
		//"-OMtemplate", project.OutputModule,
		//"-i", strconv.Itoa(project.IncrementFrame),
	}
	env := os.Environ()
	err := syscall.Exec(render.AerenderPath, args, env)
	if err != nil {
		log.Fatal(err)
	}
}
