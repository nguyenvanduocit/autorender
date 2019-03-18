package autorender

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"net/http"
	"os"
)

type Size struct {
	Width  int
	Height int
}

type Asset struct {
	Type string
	Src  string
	Name string
	Size Size
}

func (asset *Asset) DownloadTo(path string) error {
	filePath := fmt.Sprintf("%s/%s", path, asset.Name)
	output, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer output.Close()
	response, err := http.Get(asset.Src)
	if err != nil {
		return err
	}
	srcImage, _, err := image.Decode(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	dstImageFit := CoverImage(srcImage, asset.Size.Width, asset.Size.Height)
	if err := imaging.Save(dstImageFit, filePath); err != nil {
		return err
	}
	return nil
}

func CoverImage(img image.Image, width, height int) *image.NRGBA {
	maxW, maxH := width, height

	if maxW <= 0 || maxH <= 0 {
		return &image.NRGBA{}
	}

	srcBounds := img.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	if srcW <= 0 || srcH <= 0 {
		return &image.NRGBA{}
	}

	srcAspectRatio := float64(srcW) / float64(srcH)
	maxAspectRatio := float64(maxW) / float64(maxH)

	var newW, newH int
	if srcAspectRatio > maxAspectRatio {
		newW = maxW
		newH = int(float64(newW) / srcAspectRatio)
	} else {
		newH = maxH
		newW = int(float64(newH) * srcAspectRatio)
	}

	return imaging.Resize(img, newW, newH, imaging.Lanczos)
}
