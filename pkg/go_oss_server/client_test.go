package go_oss_server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)
	contentType := bodyWriter.FormDataContentType()
	const (
		ServerAddress = "http://localhost:8022"
		FileName      = "README.md"
		BucketName    = "mybucket"
		ObjectKey     = "README.md"
		AccessKey     = "root"
		AccessSecret  = "rootpw"
	)

	fileBinaryField, _ := bodyWriter.CreateFormFile("file", FileName)
	fileBin, err := os.Open(FileName)
	if err != nil {
		t.Errorf("error while open file: %s", err)
		return
	}
	defer fileBin.Close()
	io.Copy(fileBinaryField, fileBin)

	bodyWriter.WriteField("bucket", BucketName)

	bodyWriter.WriteField("object", ObjectKey)

	bodyWriter.WriteField("accesskey", AccessKey)
	bodyWriter.WriteField("secret", AccessSecret)

	bodyWriter.Close()

	req, err := http.NewRequest("POST", ServerAddress, bodyBuffer)
	req.Header.Set("Content-Type", contentType)
	client := new(http.Client)
	client.Timeout = 4 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("error while sending request: %s", err)
		return
	}
	defer resp.Body.Close()
	dat, err := ioutil.ReadAll(resp.Body)
	fmt.Println(fmt.Sprintf("resp body: %s", dat))
}
