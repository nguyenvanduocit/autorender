package autorender

import (
	"fmt"
	"github.com/xfrr/goffmpeg/transcoder"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
)

type AutoRender struct {
	AerenderPath string
}

func New(aerenderPath string) (*AutoRender, error) {
	return &AutoRender{
		AerenderPath: aerenderPath,
	}, nil
}

func (render *AutoRender) Render(project *Project) (string, error) {

	log.Println("Setup")
	if err := project.Setup(); err != nil {
		return "", err
	}

	log.Println("Prepare assets")
	if err := project.DowloadAssets(); err != nil {
		return "", err
	}

	log.Println("Replace assets")
	if err := project.ReplaceAssets(); err != nil {
		return "", err
	}

	log.Println("Start rending")
	if _, err := render.render(project); err != nil {
		return "", err
	}

	log.Println("Start convert")
	mp4Path, err := render.convert(project)
	if err != nil {
		return "", err
	}
	return mp4Path, nil
}

func (render *AutoRender) render(project *Project) (string, error) {
	movPath := path.Join(project.ProjectPath, "result.mov")
	args := []string{
		"-comp", project.Composition,
		"-project", path.Join(project.ProjectPath, project.ProjectName),
		"-output", movPath,
		// "-OMtemplate", "h264",
		"-i", strconv.Itoa(project.IncrementFrame),
	}
	cmd := exec.Command(render.AerenderPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return movPath, nil
}

func (render *AutoRender) convert(project *Project) (string, error) {

	trans := new(transcoder.Transcoder)
	mp4Path := path.Join(project.ProjectPath, "result.mp4")
	movPath := path.Join(project.ProjectPath, "result.mov")

	if err := trans.Initialize(movPath, mp4Path); err != nil {
		return "", err
	}
	done := trans.Run(true)
	progress := trans.Output()
	for msg := range progress {
		fmt.Println(msg)
	}
	if err := <-done; err != nil {
		return "", err
	}
	return mp4Path, nil
}
