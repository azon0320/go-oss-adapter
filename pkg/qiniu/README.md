# Go OSS 七牛适配器

## 七牛参数(pkg.AdapterParams)

#### 初始化需要的参数 Init()
##### `[string] conf.zone` 必须指定机房，支持以下值
* `east-china` 华东机房
* `north-china` 华北机房
* `south-china` 华南机房
* `north-usa` 北美机房
* `singapore` 新加坡机房
##### `[bool] conf.https` 是否启用HTTPS上传，默认false
##### `[bool] conf.cdn_domain` 是否启用CDN上传，默认false

#### 上传文件需要的参数
##### `[bool] policy.overwrite` 是否覆盖上传，默认false
##### `[int] policy.expires` 文件有效期，默认[七牛的设定](https://developer.qiniu.com/kodo/sdk/1238/go#5)
##### `[interface{}] upload.putret` 上传结果返回，默认为 &storage.PutRet
##### `[map[string]string] upload.extra.params` 上传自定义参数，键必须以`x:`开头，默认为空map

#### 通过字节流（Reader）上传时，必须指定的参数
##### `[int] upload.bytelen` 指定Reader读多少个字节

## 关于生成外链的参数
生成私有外链时，要指定私有链过期时间
###### `[int64] url.private.deadline` 一个unix时间戳，表示过期时间