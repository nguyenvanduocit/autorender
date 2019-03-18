package autorender

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"time"
)

type Project struct {
	TemplateFilePath string
	Composition      string
	Assets           []Asset
	ID               string
	ProjectPath      string
	ProjectName      string
	AssetPath        string
	OutputModule     string
	IncrementFrame   int
}

func NewProject(templateFilePath string, composition string, assets []Asset, outputModule string, incrementFrame int) (*Project, error) {
	if _, err := os.Stat(templateFilePath); err != nil {
		log.Fatal(fmt.Sprintf("Template is not exist: %s", templateFilePath))
	}
	fileName := path.Base(templateFilePath)
	id := fmt.Sprintf("%s--%s", fileName, time.Now().Format(time.RFC3339))
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return nil, errors.New("Can not get runtime caller")
	}
	currentDir := path.Dir(filename)

	project := Project{
		TemplateFilePath: templateFilePath,
		Composition:      composition,
		Assets:           assets,
		ID:               id,
		OutputModule:     outputModule,
		IncrementFrame:   incrementFrame,
	}
	project.ProjectName = path.Base(project.TemplateFilePath)
	project.ProjectPath = path.Join(currentDir, fmt.Sprintf("projects/%s", id))
	project.AssetPath = path.Join(project.ProjectPath, "assets")
	return &project, nil
}

func (project *Project) Setup() error {
	if err := os.MkdirAll(project.ProjectPath, 0777); err != nil {
		return err
	}
	if err := os.MkdirAll(project.AssetPath, 0777); err != nil {
		return err
	}
	return nil
}

func (project *Project) DowloadAssets() error {
	for _, asset := range project.Assets {
		if err := asset.DownloadTo(project.AssetPath); err != nil {
			return err
		}
	}
	return nil
}

/**
 * Currently, It can only replace filepath, Which can be different from what displaced in AE UI. Because in AE UI, you can rename asset, but the real filename not change.
 */
func (project *Project) ReplaceAssets() error {
	contentBytes, err := ioutil.ReadFile(project.TemplateFilePath)
	if err != nil {
		return err
	}
	content := string(contentBytes)
	for _, asset := range project.Assets {
		content = regexp.MustCompile(fmt.Sprintf(`fullpath="([^"]*%s[^"]*)"`, asset.Name)).ReplaceAllString(content, fmt.Sprintf(`fullpath="%s"`, path.Join("", path.Join(project.AssetPath, asset.Name))))
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", project.ProjectPath, project.ProjectName), []byte(content), 0777)
	if err != nil {
		return err
	}
	return nil
}

func (project *Project) ClearTempDir() {
	//os.RemoveAll(project.ProjectPath)
}
