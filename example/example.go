package main

import (
	"github.com/nguyenvanduocit/autorender"
	"log"
)

func main() {
	render, err := autorender.New("/Applications/Adobe After Effects CC 2019/aerender")
	if err != nil {
		log.Fatal(err)
	}
	assets := []autorender.Asset{
		{
			Type: "image",
			Src:  "http://www.stickpng.com/assets/images/58482ec0cef1014c0b5e4a70.png",
			Name: "Adobe_After_Effects_CC_Logo.png",
			Size: autorender.Size{
				Width:  2000,
				Height: 1950,
			},
		},
	}
	project, err := autorender.NewProject("/Users/duocnguyen/autorender/templates/template1/template.aepx", "Main Comp", assets, "h264", 1)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Project created: ", project.ID)
	mp4Path, err := render.Render(project)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("MP4: %s\n", mp4Path)
}
