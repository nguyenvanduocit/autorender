package autorender

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
)

type Project struct {
	TemplateFilePath string
	Composition string
	Assets []Asset
	ID string
	ProjectPath string
	ProjectName string
	AssetPath string
	OutputModule string
	IncrementFrame int
}

func NewProject(templateFilePath string, composition string, assets []Asset, outputModule string, incrementFrame int)(*Project){
	id, err := pseudo_uuid()
	if err != nil {
		log.Fatal(err)
	}
	_, filename, _, ok := runtime.Caller(1)

	if(!ok){
		log.Fatal("Can not get runtime caller")
	}
	currentDir := path.Dir(filename)

	project := Project{
		TemplateFilePath: templateFilePath,
		Composition: composition,
		Assets: assets,
		ID: id,
		OutputModule: outputModule,
		IncrementFrame: incrementFrame,
	}
	project.ProjectName = path.Base(project.TemplateFilePath)
	project.ProjectPath = path.Join(currentDir, fmt.Sprintf("projects/%s", id))
	project.AssetPath = path.Join(project.ProjectPath, "assets")
	return &project
}

func (project *Project)Validate(){
	if _, err := os.Stat(project.TemplateFilePath); err != nil {
		log.Fatal(fmt.Sprintf("Template is not exist: %s", project.TemplateFilePath))
	}
}

// Setup enviroment:
// 1. Create directory
// 2. Copy template
func (project *Project)Setup(){
	if err := os.MkdirAll(project.ProjectPath, 0777); err != nil{
		log.Fatal(err)
	}
	if err := os.MkdirAll(project.AssetPath, 0777); err != nil{
		log.Fatal(err)
	}
}

// TODO: Use routime
func (project *Project)DowloadAssets(){
	for _, asset := range project.Assets{
		asset.DownloadTo(project.AssetPath)
	}
	log.Println("All image downloadee")
}

/**
 * Currently, It can only replace filepath, Which can be different from what displaced in AE UI. Because in AE UI, you can rename asset, but the real filename not change.
 */
func (project *Project)ReplaceAssets(){
	fmt.Println("Replace Assets")
	contentBytes, err := ioutil.ReadFile(project.TemplateFilePath)
	if err != nil {
		log.Fatal(err)
	}
	content := string(contentBytes)
	for _, asset := range project.Assets{
		content = regexp.MustCompile(fmt.Sprintf(`fullpath="([^"]*%s[^"]*)"`, asset.Name)).ReplaceAllString(content, fmt.Sprintf(`fullpath="%s"`,path.Join("", path.Join(project.AssetPath, asset.Name))))
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", project.ProjectPath, project.ProjectName), []byte(content), 0777)
	if err != nil {
		panic(err)
	}
}

func (project *Project)ClearTempDir(){
	//os.RemoveAll(project.ProjectPath)
}
