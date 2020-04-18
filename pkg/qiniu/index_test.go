package qiniu

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"testing"
)

func TestUpload(t *testing.T) {
	accessKey := "myaccesskey"
	accessSecret := "myaccesssecret"
	objectKey := "1.png"
	localFile := "1.png"
	bucket := "mybucket"
	overwriteMode := false
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	if overwriteMode {
		putPolicy.Scope = fmt.Sprintf("%s:%s", bucket, objectKey)
	}
	mac := qbox.NewMac(accessKey, accessSecret)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}

	_ = formUploader.PutFile(context.Background(), &ret, upToken, objectKey, localFile, &putExtra)
}
