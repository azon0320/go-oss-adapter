package go_oss_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dormao/go-oss-adapter/pkg"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	AdapterName = "go-oss-server"
)

type BaseResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

type Adapter struct {
	credentials *pkg.CredentialsToken
	bucket      *string
	EndpointURL *string
	client      *http.Client
}

func (adapter *Adapter) Init(token pkg.CredentialsToken, params pkg.AdapterParams) error {
	adapter.EndpointURL = &token.Endpoint
	adapter.credentials = &token
	var timeout = int64(params.GetOrDefault(ParamKeyTimeout, 3).(int))
	if timeout == 0 {
		timeout = 3
	}
	adapter.client = &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	clint := &http.Client{}
	req, _ := http.NewRequest("OPTIONS", *adapter.EndpointURL, bytes.NewReader([]byte("hello")))
	resp, err := clint.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("checking connection for go-oss-server (%s) failed", *adapter.EndpointURL))
	}
	return nil
}

func (adapter *Adapter) Bucket(buck string) error {
	adapter.bucket = &buck
	return nil
}

func (adapter *Adapter) GetBucket() *string {
	return adapter.bucket
}

func (adapter *Adapter) Name() string { return AdapterName }

func (adapter *Adapter) PutObjectFromByteArray(key string, data []byte, readLen int64, params pkg.AdapterParams) (interface{}, error) {
	if adapter.bucket == nil {
		return nil, errors.New(fmt.Sprintf("bucket not specified"))
	}
	reader := bytes.NewReader(data)
	return adapter.PutObjectFromReader(key, reader, params)
}

func (adapter *Adapter) PutObjectFromReader(key string, reader io.Reader, params pkg.AdapterParams) (interface{}, error) {
	if adapter.bucket == nil {
		return nil, errors.New(fmt.Sprintf("bucket not specified"))
	}
	ext := params.GetOrDefault(ParamKeyExt, "txt").(string)
	req, err := adapter.buildUploadForm(adapter.credentials, *adapter.bucket, key, reader, fmt.Sprintf("%s.%s", "file", ext))
	if err != nil {
		return nil, err
	}
	resp, err := adapter.client.Do(req)
	if err != nil {
		return nil, err
	}
	return adapter.onResponse(resp)
}

func (adapter *Adapter) PutObjectFromFilePath(key, filepath string, params pkg.AdapterParams) (interface{}, error) {
	if adapter.bucket == nil {
		return nil, errors.New(fmt.Sprintf("bucket not specified"))
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	req, err := adapter.buildUploadForm(adapter.credentials, *adapter.bucket, key, file, file.Name())
	if err != nil {
		return nil, err
	}
	resp, err := adapter.client.Do(req)
	if err != nil {
		return nil, err
	}
	return adapter.onResponse(resp)
}

func (adapter *Adapter) GetObjectToClosableReader(key string, params pkg.AdapterParams) (io.ReadCloser, error) {
	return nil, errors.New(fmt.Sprintf("get object to io.Reader is not supported by %s", adapter.Name()))
}

func (adapter *Adapter) GetObjectToBytes(key string, params pkg.AdapterParams) ([]byte, error) {
	return nil, errors.New(fmt.Sprintf("get object to bytes is not supported by %s", adapter.Name()))
}

func (adapter *Adapter) GetObjectToFile(key string, params pkg.AdapterParams) error {
	return errors.New(fmt.Sprintf("get object to file is not supported by %s", adapter.Name()))
}

func (adapter *Adapter) MakePublicURL(key string, params pkg.AdapterParams) string {
	if adapter.bucket == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", *adapter.EndpointURL, *adapter.bucket, strings.Trim(key, "/"))
}

func (adapter *Adapter) MakePrivateURL(key string, params pkg.AdapterParams) string {
	if adapter.bucket == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", *adapter.EndpointURL, *adapter.bucket, strings.Trim(key, "/"))
}

func (adapter *Adapter) ListObjects(keyPrefix string, params pkg.AdapterParams) (list []string, err error) {
	if adapter.bucket == nil {
		return []string{}, errors.New(fmt.Sprintf("bucket not specified"))
	}
	const (
		Method             = "POST"
		FormFieldAccessKey = "accesskey"
		FormFieldSecret    = "secret"
		FormFieldBucket    = "bucket"
	)
	var buffer = new(bytes.Buffer)
	var writer = multipart.NewWriter(buffer)
	list = make([]string, 0)
	err = nil
	contentType := writer.FormDataContentType()
	err = writer.WriteField(FormFieldAccessKey, adapter.credentials.AccessKey)
	if err != nil {
		return
	}
	err = writer.WriteField(FormFieldSecret, adapter.credentials.AccessSecret)
	if err != nil {
		return
	}
	err = writer.WriteField(FormFieldBucket, *adapter.bucket)
	if err != nil {
		return
	}
	writer.Close()
	req, err := http.NewRequest(Method, fmt.Sprintf("%s?%s=%s", *adapter.EndpointURL, "list", keyPrefix), buffer)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := adapter.client.Do(req)
	if err != nil {
		return
	}
	baseResp, err := adapter.onResponse(resp)
	if err != nil {
		return
	}
	var mapped = baseResp.Result.(map[string]interface{})
	for _, v := range mapped["result"].([]interface{}) {
		list = append(list, strings.Trim(fmt.Sprintln(v), "\n"))
	}
	return
}

func (adapter *Adapter) DeleteObject(key string, params pkg.AdapterParams) (v interface{}, err error) {
	if adapter.bucket == nil {
		return []string{}, errors.New(fmt.Sprintf("bucket not specified"))
	}
	const (
		Method             = "DELETE"
		FormFieldAccessKey = "accesskey"
		FormFieldSecret    = "secret"
		FormFieldBucket    = "bucket"
		FormFieldObject    = "object"
	)
	var buffer = new(bytes.Buffer)
	var writer = multipart.NewWriter(buffer)
	v = nil
	err = nil
	contentType := writer.FormDataContentType()
	err = writer.WriteField(FormFieldAccessKey, adapter.credentials.AccessKey)
	if err != nil {
		return
	}
	err = writer.WriteField(FormFieldSecret, adapter.credentials.AccessSecret)
	if err != nil {
		return
	}
	err = writer.WriteField(FormFieldBucket, *adapter.bucket)
	if err != nil {
		return
	}
	err = writer.WriteField(FormFieldObject, key)
	if err != nil {
		return
	}
	writer.Close()
	req, err := http.NewRequest(Method, *adapter.EndpointURL, buffer)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := adapter.client.Do(req)
	if err != nil {
		return
	}
	v, err = adapter.onResponse(resp)
	return
}

func (adapter *Adapter) GetUploadToken(params pkg.AdapterParams) (tkn interface{}, err error) {
	tkn = ""
	return
}

func (adapter *Adapter) buildUploadForm(
	creden *pkg.CredentialsToken,
	bucket string, key string,
	reader io.Reader, filename string,
) (req *http.Request, err error) {
	const (
		Method             = "POST"
		FormFieldFile      = "file"
		FormFieldAccessKey = "accesskey"
		FormFieldSecret    = "secret"
		FormFieldBucket    = "bucket"
		FormFieldObject    = "object"
	)
	var buffer = new(bytes.Buffer)
	var writer = multipart.NewWriter(buffer)
	contentType := writer.FormDataContentType()
	fileFieldWriter, err := writer.CreateFormFile(FormFieldFile, filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileFieldWriter, reader)
	if err != nil {
		return nil, err
	}
	err = writer.WriteField(FormFieldAccessKey, creden.AccessKey)
	if err != nil {
		return nil, err
	}
	err = writer.WriteField(FormFieldSecret, creden.AccessSecret)
	if err != nil {
		return nil, err
	}
	err = writer.WriteField(FormFieldBucket, bucket)
	if err != nil {
		return nil, err
	}
	err = writer.WriteField(FormFieldObject, key)
	if err != nil {
		return nil, err
	}
	writer.Close()
	req, err = http.NewRequest(Method, *adapter.EndpointURL, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return
}

func (adapter *Adapter) onResponse(resp *http.Response) (*BaseResponse, error) {
	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var baseResp = BaseResponse{}
	json.Unmarshal(datas, &baseResp)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("go-oss-server returned an non-200 status (code: %d, msg: %s)", baseResp.Code, baseResp.Msg))
	}
	return &baseResp, nil
}
