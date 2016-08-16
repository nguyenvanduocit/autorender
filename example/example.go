package main

import (
	"github.com/nguyenvanduocit/autorender"
	"log"
)

func main() {
	render,err := autorender.New("/Applications/Adobe After Effects CC 2015.3/aerender")
	if err != nil {
		log.Fatal(err)
	}
	assets := []autorender.Asset{
		{
				Type: "image",
				Src: "http://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
				Name: "pokemon.png",
			},
	}
	project := autorender.NewProject("/Users/duoc/.go/src/github.com/nguyenvanduocit/autorender/example/templates/dragonite.aepx", "main", assets, "ahihi", 1)
	log.Println("Project created: ", project.ID)
	render.Render(project)
}
