package controllers

import (
	"alfath_lms/api/funcs"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	UploadHandler struct {
		responder *web.Responder
	}
)

func (uploadHandler *UploadHandler) Inject(
	responder *web.Responder,
) {
	uploadHandler.responder = responder
}

func (uploadHandler *UploadHandler) Setup(ctx context.Context, req *web.Request) *web.Response {
	if req.Params["file_name"] == "" {
		return funcs.CorsedResponse(uploadHandler.responder.HTTP(400, strings.NewReader("Please select a file")))
	}

	fmt.Println(req.Params["file_name"])

	filePath := "./uploads"
	filePath = filepath.Join(filePath, req.Params["file_name"])
	file, err := os.Open(filePath)
	if err != nil {
		return funcs.CorsedResponse(uploadHandler.responder.HTTP(400, strings.NewReader("File doesn't exist")))
	}

	fileReader := io.Reader(file)

	fileSplit := strings.Split(req.Params["file_name"], ".")
	fileType := ""
	if fileSplit[1] == "jpg" || fileSplit[1] == "png" || fileSplit[1] == "jpeg" {
		fileType = "image/" + fileSplit[1]
	}

	resp := funcs.CorsedResponse(uploadHandler.responder.Download(fileReader, fileType, req.Params["file_name"], false))
	//defer file.Close()
	return resp
}
