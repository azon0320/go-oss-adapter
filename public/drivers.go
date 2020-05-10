package public

import (
	"github.com/azon0320/go-oss-adapter/internal/go_oss_server"
	"github.com/azon0320/go-oss-adapter/internal/qiniu"
	"github.com/azon0320/go-oss-adapter/pkg"
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
