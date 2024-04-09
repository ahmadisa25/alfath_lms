package controllers

import (
	"alfath_lms/api/funcs"
	"context"
	"fmt"
	"net/http"
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

	/*currentDir, err := os.Getwd()
	if err != nil {
		return funcs.CorsedResponse(uploadHandler.responder.HTTP(500, strings.NewReader("Internal server error")))
	}
	filePath := filepath.Join(currentDir, "../uploads")*/
	filePath := "./uploads"
	filePath = filepath.Join(filePath, req.Params["file_name"])
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return funcs.CorsedResponse(uploadHandler.responder.HTTP(400, strings.NewReader("File doesn't exist")))
	}

	fileType := ""
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return funcs.CorsedResponse(uploadHandler.responder.HTTP(400, strings.NewReader("File read error")))
	}

	// Determine the content type
	fileType = http.DetectContentType(buffer)

	responseHeader := make(http.Header)
	responseHeader.Set("Content-Type", fileType)
	responseHeader.Set("Content-Disposition", "inline")

	//need to reset the pointer back to the beginning after reading the file.
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return funcs.CorsedResponse(uploadHandler.responder.HTTP(500, strings.NewReader("File seek error")))
	}

	resp := &web.Response{
		Status:         http.StatusOK,
		Body:           file, //don't need to close the file because if we used Body in http response, the file should be closed automatically
		Header:         responseHeader,
		CacheDirective: nil, // You may set cache directives if needed
	}

	return resp

}
