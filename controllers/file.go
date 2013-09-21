// file
package controllers

import (
	"fmt"
	"github.com/ginuerzh/icecar/errors"
	"github.com/ginuerzh/weedo"
	"io"
	"log"
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

	log.Println(header.Filename)
	fid, size, err := weedo.AssignUpload(header.Filename, file)
	if err != nil {
		fmt.Println(err)
		this.Data["json"] = this.response(nil, &errors.FileUploadError)
		this.ServeJson()
		return
	}
	resp := map[string]interface{}{"fid": fid, "name": header.Filename, "size": size}
	this.Data["json"] = this.response(resp, nil)
	this.ServeJson()
}

func (this *FileController) Download() {
	fid := this.Ctx.Input.Param[":all"]
	file, err := weedo.Download(fid)
	if err != nil {
		this.Data["json"] = this.response(nil, &errors.FileNotFoundError)
		this.ServeJson()
		return
	}
	//url, _ := weedo.GetUrl(fid)
	//this.Redirect(url, 302)
	defer file.Close()

	io.Copy(this.Ctx.ResponseWriter, file)
}

func (this *FileController) Delete() {
	fid := this.Ctx.Input.Param[":all"]
	if err := weedo.Delete(fid); err != nil {
		log.Println(err)
	}

	this.Data["json"] = this.response(nil, nil)
	this.ServeJson()
}
