package main

import (
	"context"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/vidsvr/viddirect"
)

func main() {
	defer utils.Recover(context.Background())
	svcCtx := viddirect.GetSvcCtx()
	viddirect.ApiDirectRun()
	viddirect.Run(svcCtx)
}
