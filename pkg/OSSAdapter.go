package pkg

import "io"

type OSSAdapter interface {
	Init(token CredentialsToken, params AdapterParams) error
	Bucket(buck string) error
	Name() string

	PutObjectFromByteArray(key string, data []byte, readLen int64, params AdapterParams) (interface{}, error)
	PutObjectFromReader(key string, reader io.Reader, params AdapterParams) (interface{}, error)
	PutObjectFromFilePath(key, filepath string, params AdapterParams) (interface{}, error)

	GetObjectToClosableReader(key string, params AdapterParams) (io.ReadCloser, error)
	GetObjectToBytes(key string, params AdapterParams) ([]byte, error)
	GetObjectToFile(key string, params AdapterParams) error

	MakePublicURL(key string, params AdapterParams) string
	MakePrivateURL(key string, params AdapterParams) string
}
