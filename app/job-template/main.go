package main

import (
	_ "my-crontab/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"my-crontab/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
