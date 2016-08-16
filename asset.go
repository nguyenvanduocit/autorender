package autorender

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

type Asset struct{
	Type string
	Src string
	Name string
}

func (asset *Asset)DownloadTo(path string)(int64, error){
	filePath := fmt.Sprintf("%s/%s", path, asset.Name)
	output, err := os.Create(filePath)
	if err != nil {
		return -1,err
	}
	defer output.Close()
	response, err := http.Get(asset.Src)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()
	n, err := io.Copy(output, response.Body)
	if err != nil {
		return -1, err
	}
	return n, nil
}
