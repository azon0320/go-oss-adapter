package qiniu

import "github.com/dormao/go-oss-adapter/pkg"

const (
	// storage conf , used from Init
	ParamKeyZone          = pkg.ParamKeyZone          // string
	ParamKeyUseHTTPS      = pkg.ParamKeyUseHTTPS      // bool
	ParamKeyUseCDNDomains = pkg.ParamKeyUseCDNDomains // bool
)

const (
	// upload without overwrite:  ${BUCKET_NAME}
	// upload with overwrite: ${BUCKET_NAME}:${KEY_NAME}
	ParamKeyPolicyOverwrite = pkg.ParamKeyPolicyOverwrite // bool
	ParamKeyPolicyExpires   = pkg.ParamKeyPolicyExpires   // uint64 unit:second
	ParamKeyPolicyObject    = pkg.ParamKeyPolicyObject    // string
	ParamKeyListLimit       = pkg.ParamKeyListLimit       // int
)

const (
	ParamKeyByteLen     = pkg.ParamKeyByteLen     // int
	ParamKeyPutRet      = pkg.ParamKeyPutRet      // struct pointer
	ParamKeyExtraParams = pkg.ParamKeyExtraParams // map[string]string
)

const (
	ParamKeyPrivateURLDeadlineUnix = pkg.ParamKeyPrivateURLDeadlineUnix
)
