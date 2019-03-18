package main

import (
	"github.com/xfrr/goffmpeg/transcoder"
	"gopkg.in/cheggaaa/pb.v2"
	"log"
)

func main() {
	trans := new(transcoder.Transcoder)
	mp4Path := "/Users/duocnguyen/go/src/github.com/nguyenvanduocit/autorender/example/projects/template.aepx--2019-03-18T15:23:57+07:00/result.mp4"
	movPath := "/Users/duocnguyen/go/src/github.com/nguyenvanduocit/autorender/example/projects/template.aepx--2019-03-18T15:23:57+07:00/result.mov"

	if err := trans.Initialize(movPath, mp4Path); err != nil {
		log.Panicln(err)
	}
	done := trans.Run(true)
	progress := trans.Output()
	bar := pb.StartNew(100)
	for msg := range progress {
		bar.Add(int(msg.Progress))
	}
	if err := <-done; err != nil {
		log.Panicln(err)
	}
}
