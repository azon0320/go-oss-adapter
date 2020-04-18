package pkg

type CredentialsToken struct {
	AccessKey    string
	AccessSecret string
	Endpoint     string
}

type AdapterParams map[string]interface{}

func (params AdapterParams) Get(key string) (exists bool, value interface{}) {
	value, exists = params[key]
	return
}

func (params AdapterParams) GetOrDefault(key string, def interface{}) interface{} {
	ex, v := params.Get(key)
	if ex {
		return v
	} else {
		return def
	}
}
