// file
package controllers

import (
	"fmt"
	"github.com/ginuerzh/icecar/errors"
	"github.com/ginuerzh/weedo"
)

type FileController struct {
	BaseController
}

func (this *FileController) Upload() {
	file, header, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = this.response(nil, &errors.FileNotFoundError)
		this.ServeJson()
		return
	}

	client := weedo.NewClient("localhost:9333")
	fid, err := client.Upload(header.Filename, file)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = this.response(nil, &errors.FileUploadError)
		this.ServeJson()
		return
	}
	url, err := client.GetUrl(fid)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = this.response(nil, &errors.FileUploadError)
		this.ServeJson()
		return
	}

	this.Data["json"] = this.response(map[string]interface{}{"url": url}, nil)
	this.ServeJson()
}
