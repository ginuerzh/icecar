// file
package controllers

import (
	"fmt"
	"github.com/ginuerzh/weedo"
	"icecar/errors"
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

	fmt.Println(string(this.Ctx.RequestBody))

	client := weedo.NewClient("http://localhost:9333")
	url, err := client.Upload(header.Filename, header.Header.Get("Content-Type"), file)
	if err != nil {
		fmt.Println(err)
	}
	this.Data["json"] = this.response(map[string]interface{}{"url": url}, nil)
	this.ServeJson()
}
