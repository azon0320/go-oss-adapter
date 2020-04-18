package qiniu

const (
	// storage conf , used from Init
	ParamKeyZone          = "conf.zone"       // string
	ParamKeyUseHTTPS      = "conf.https"      // bool
	ParamKeyUseCDNDomains = "conf.cdn_domain" // bool
)

const (
	// upload without overwrite:  ${BUCKET_NAME}
	// upload with overwrite: ${BUCKET_NAME}:${KEY_NAME}
	ParamKeyPolicyOverwrite = "policy.overwrite" // bool
	ParamKeyPolicyExpires   = "policy.expires"   // uint64 unit:second
)

const (
	ParamKeyByteLen     = "upload.bytelen"      // int
	ParamKeyPutRet      = "upload.putret"       // struct pointer
	ParamKeyExtraParams = "upload.extra.params" // map[string]string
)

const (
	ParamKeyPrivateURLDeadlineUnix = "url.private.deadline"
)
