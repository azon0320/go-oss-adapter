package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/dormao/go-oss-adapter/pkg"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"io"
	"time"
)

const (
	AdapterName = "qiniu"
)

type Adapter struct {
	mac      *qbox.Mac
	conf     *storage.Config
	uploader *storage.FormUploader
	bucket   *string
	Domain   *string
}

func (adapter *Adapter) Init(token pkg.CredentialsToken, params pkg.AdapterParams) error {
	adapter.Domain = &token.Endpoint
	adapter.mac = qbox.NewMac(token.AccessKey, token.AccessSecret)
	var zone = GetZoneFromString(params.GetOrDefault(ParamKeyZone, ZoneHuadong).(string))
	var useHttps = params.GetOrDefault(ParamKeyUseHTTPS, false).(bool)
	var usecdn = params.GetOrDefault(ParamKeyUseCDNDomains, false).(bool)
	adapter.conf = &storage.Config{
		Zone: zone, UseHTTPS: useHttps, UseCdnDomains: usecdn,
	}
	adapter.uploader = storage.NewFormUploader(adapter.conf)
	return nil
}

func (adapter *Adapter) Bucket(buck string) error {
	adapter.bucket = &buck
	return nil
}

func (adapter *Adapter) Name() string { return AdapterName }

func (adapter *Adapter) PutObjectFromByteArray(key string, data []byte, readLen int64, params pkg.AdapterParams) (interface{}, error) {
	var reader = bytes.NewReader(data)
	upToken, _, extra, ret := adapter.prepareUploadEssentials(key, params)
	err := adapter.uploader.Put(context.Background(), ret, upToken, key, reader, readLen, extra)
	return ret, err
}

func (adapter *Adapter) PutObjectFromReader(key string, reader io.Reader, params pkg.AdapterParams) (interface{}, error) {
	var readLen = int64(params.GetOrDefault(ParamKeyByteLen, 0).(int))
	if readLen <= 0 {
		return nil, errors.New(fmt.Sprintf("specified param key (%s) not found", ParamKeyByteLen))
	}
	upToken, _, extra, ret := adapter.prepareUploadEssentials(key, params)
	err := adapter.uploader.Put(context.Background(), ret, upToken, key, reader, readLen, extra)
	return ret, err
}

func (adapter *Adapter) PutObjectFromFilePath(key, filepath string, params pkg.AdapterParams) (interface{}, error) {
	upToken, _, extra, ret := adapter.prepareUploadEssentials(key, params)
	err := adapter.uploader.PutFile(context.Background(), ret, upToken, key, filepath, extra)
	return ret, err
}

func (adapter *Adapter) GetObjectToClosableReader(key string, params pkg.AdapterParams) (io.ReadCloser, error) {
	return nil, errors.New(fmt.Sprintf("get object to io.Reader is not supported by %s", adapter.Name()))
}

func (adapter *Adapter) GetObjectToBytes(key string, params pkg.AdapterParams) ([]byte, error) {
	return nil, errors.New(fmt.Sprintf("get object to bytes[] is not supported by %s", adapter.Name()))
}

func (adapter *Adapter) GetObjectToFile(key string, params pkg.AdapterParams) error {
	return errors.New(fmt.Sprintf("get object to file is not supported by %s, use %s instead", adapter.Name(), "URL Download"))
}

func (adapter *Adapter) MakePublicURL(key string, params pkg.AdapterParams) string {
	return storage.MakePublicURL(*adapter.Domain, key)
}

func (adapter *Adapter) MakePrivateURL(key string, params pkg.AdapterParams) string {
	return storage.MakePrivateURL(
		adapter.mac,
		*adapter.Domain, key,
		params.GetOrDefault(
			ParamKeyPrivateURLDeadlineUnix, time.Now().Add(5*time.Minute).Unix()).(int64),
	)
}

func (adapter *Adapter) prepareUploadEssentials(key string, params pkg.AdapterParams) (upToken string, putPolicy *storage.PutPolicy, putExtra *storage.PutExtra, putRet interface{}) {
	putPolicy = &storage.PutPolicy{}
	overWriteMode := params.GetOrDefault(ParamKeyPolicyOverwrite, false).(bool)
	if overWriteMode {
		putPolicy.Scope = fmt.Sprintf("%s:%s", *adapter.bucket, key)
	} else {
		putPolicy.Scope = *adapter.bucket
	}
	expires := params.GetOrDefault(ParamKeyPolicyExpires, 0).(int)
	if expires > 0 {
		putPolicy.Expires = uint64(expires)
	}
	putRet = params.GetOrDefault(ParamKeyPutRet, nil)
	if putRet == nil {
		putRet = &storage.PutRet{}
	}
	putExtra = &storage.PutExtra{}
	putExtra.Params = params.GetOrDefault(ParamKeyExtraParams, map[string]string{}).(map[string]string)
	upToken = putPolicy.UploadToken(adapter.mac)
	return
}
