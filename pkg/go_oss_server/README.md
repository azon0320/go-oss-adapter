# Go OSS go-oss-server适配器

## go-oss-server的针对性参数(pkg.AdapterParams)

#### 初始化需要的参数 Init()
##### `[int] conf.timeout` 超时时间，单位（秒），默认3

#### 上传非文件需要的参数
##### `[string] upload.ext` 文件扩展名，建议指定，这样可以帮助浏览器识别文件类型，默认txt

## 示例用法
~~~
# 要上传的文件和对象key
const (
    AccessKey = "myaccess_key"
    AccessSecret = "myaccess_secret"
    AccessBucket = "mybucket"

    # go-oss-server 的 endpoint 一定要填上 http 或 https
    OSSServerEndpoint = "http://localhost"

    Key = "file.exe"
    File = "./file.exe"
)

# 初始化
var ossClient = make(go_oss_server.Adapter)
err := ossClient.Init(pkg.CredentialsToken{
	AccessKey:    AccessKey,
	AccessSecret: AccessSecret,
	Endpoint:     OSSServerEndpoint,
}, pkg.AdapterParams{})

# 指定桶
ossClient.Bucket(AccessBucket)

# 针对性参数
var params = pkg.AdapterParams{}

# 上传本地文件
resp, err := ossClient.PutObjectFromFilePath(Key, File, params)

# 上传字节数组
var ReadLen = 0 // go-oss-server 不支持限制Reader的读取长度
resp, err := ossClient.PutObjectFromByteArray(Key, []byte("Hello I am poweredormao"), ReadLen, params)

# 获得公链(go-oss-server不支持私有链)
ossClient.MakePublicURL(Key, params)
~~~