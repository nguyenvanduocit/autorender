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
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/12371208_10201352346168296_7908416656263087848_o.jpg?dl=1",
			Name: "1.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/12375334_10201352344568256_4787472564867341244_o.jpg?dl=1",
			Name: "2.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/921224_10201352348848363_1788953854725095452_o.jpg?dl=1",
			Name: "3.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/12362930_10201352346648308_7084547198569257564_o.jpg?dl=1",
			Name: "4.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/12371219_10201352356608557_7367004372475008160_o.jpg?dl=1",
			Name: "5.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/12006635_10205545911461842_9072372145096976564_o.jpg?dl=1",
			Name: "6.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/1622511_244124209091745_815460841_o.jpg?dl=1",
			Name: "7.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/1262460_4719031391969_856951380_o.jpg?dl=1",
			Name: "8.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/1102705_661343873883602_1300801235_o.jpg?dl=1",
			Name: "9.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/13653264_10204903338541460_3091639653897474252_o.jpg?dl=1",
			Name: "10.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/13581985_10204893600738021_6553305135435591288_o.jpg?dl=1",
			Name: "11.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/13576893_10204877047844209_2228288144167103565_o.jpg?dl=1",
			Name: "12.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/13603573_10204876538031464_6021401051227477878_o.jpg?dl=1",
			Name: "13.jpg",
		},{
			Type: "image",
			Src: "https://scontent-hkg3-1.xx.fbcdn.net/t31.0-8/13502891_10204820800678065_6500548131696542679_o.jpg?dl=1",
			Name: "14.jpg",
		},
	}
	project := autorender.NewProject("/Users/duoc/.go/src/github.com/nguyenvanduocit/autorender/example/templates/dragonite.aepx", "Render", assets, "ahihi", 1)
	log.Println("Project created: ", project.ID)
	render.Render(project)
}
