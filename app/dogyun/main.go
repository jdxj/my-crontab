package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"my-crontab/app/dogyun/internal/cmd"
	_ "my-crontab/app/dogyun/internal/logic"
	_ "my-crontab/app/dogyun/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
