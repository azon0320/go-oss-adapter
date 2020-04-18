package go_oss_server

import (
	"encoding/json"
	"fmt"
	"github.com/dormao/go-oss-adapter/pkg"
	"testing"
)

func TestUploadFile(t *testing.T) {
	var ossClient = Adapter{}
	err := ossClient.Init(pkg.CredentialsToken{
		AccessKey:    "root",
		AccessSecret: "rootpw",
		Endpoint:     "http://localhost:8022",
	}, pkg.AdapterParams{})
	ossClient.Bucket("mybucket")
	if err != nil {
		t.Errorf("error init go-oss-server client: %s", err)
		return
	}
	resp, err := ossClient.PutObjectFromFilePath("dormao.md", "README.md", pkg.AdapterParams{})
	if err != nil {
		t.Errorf("error put object to oss: %s", err)
		return
	}
	respStr, _ := json.Marshal(resp)
	fmt.Println(fmt.Sprintf("BaseResponse: %s", string(respStr)))
	fmt.Println(fmt.Sprintf("generate public URL: %s", ossClient.MakePublicURL("README.md", pkg.AdapterParams{})))
}

func TestExecutable(t *testing.T) {
	const (
		Key  = "keyboard_test.exe"
		File = "./keyboard_test.exe"
	)
	var ossClient = Adapter{}
	err := ossClient.Init(pkg.CredentialsToken{
		AccessKey:    "root",
		AccessSecret: "rootpw",
		Endpoint:     "http://localhost:8022",
	}, pkg.AdapterParams{})
	ossClient.Bucket("mybucket")
	if err != nil {
		t.Errorf("error init go-oss-server client: %s", err)
		return
	}
	resp, err := ossClient.PutObjectFromFilePath(Key, File, pkg.AdapterParams{})
	if err != nil {
		t.Errorf("error put object to oss: %s", err)
		return
	}
	respStr, _ := json.Marshal(resp)
	fmt.Println(fmt.Sprintf("BaseResponse: %s", string(respStr)))
	fmt.Println(fmt.Sprintf("generate public URL: %s", ossClient.MakePublicURL(Key, pkg.AdapterParams{})))
}

func TestBytes(t *testing.T) {
	const (
		Key = "dormao"
	)
	var ossClient = Adapter{}
	err := ossClient.Init(pkg.CredentialsToken{
		AccessKey:    "root",
		AccessSecret: "rootpw",
		Endpoint:     "http://localhost:8022",
	}, pkg.AdapterParams{})
	ossClient.Bucket("mybucket")
	resp, err := ossClient.PutObjectFromByteArray(Key, []byte("Hello i am dormao"), 0, pkg.AdapterParams{})
	if err != nil {
		t.Errorf("error put object to oss: %s", err)
		return
	}
	respStr, _ := json.Marshal(resp)
	fmt.Println(fmt.Sprintf("BaseResponse: %s", string(respStr)))
	fmt.Println(fmt.Sprintf("generate public URL: %s", ossClient.MakePublicURL(Key, pkg.AdapterParams{})))
}
