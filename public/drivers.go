package public

import (
	"gitlab.linkjob.top/luoyongjin/go-oss-adapter/internal/go_oss_server"
	"gitlab.linkjob.top/luoyongjin/go-oss-adapter/internal/qiniu"
	"gitlab.linkjob.top/luoyongjin/go-oss-adapter/pkg"
)

func CreateAdapter(adaptername string) pkg.OSSAdapter {
	switch adaptername {
	case "qiniu":
		return &qiniu.Adapter{}
	case "go-oss-server":
		return &go_oss_server.Adapter{}
	default:
		return nil
	}
}
