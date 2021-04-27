package models

import (
	"awesomeProject3/utils/errmsg"
	"net/http"
	"os"
	"io"
	"bytes"
	"mime/multipart"
)

func UpLoadFile(url, file string)int {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return errmsg.ERROR
	}
	defer f.Close()
	fw, err := w.CreateFormFile("file", file)
	if err != nil {
		return errmsg.ERROR
	}
	if _, err = io.Copy(fw, f); err != nil {
		return errmsg.ERROR
	}
	// Add the other fields
	if fw, err = w.CreateFormField("key"); err != nil {
		return errmsg.ERROR
	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		return errmsg.ERROR
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return errmsg.ERROR
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return errmsg.ERROR
	}
	// Check the response
	if res.StatusCode != http.StatusOK {
		return errmsg.ERROR
	}
	return errmsg.SUCCESE
}