package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gproc"

	"my-crontab/app/dogyun/internal/model"
	"my-crontab/app/dogyun/internal/service"
)

var (
	Main = gcmd.Command{
		Name:   "dogyun",
		Usage:  "start crontab job",
		Strict: true,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, `cron job start`)

			every, err := g.Cfg().Get(ctx, "dogyun.every")
			if err != nil {
				return err
			}

			pattern := fmt.Sprintf("@every %s", every.String())
			_, err = gcron.Add(ctx, pattern, func(ctx context.Context) {
				changed, err := service.Content().GetChangedProducts(ctx)
				if err != nil {
					g.Log().Errorf(ctx, "get changed products err: %s", err)
					return
				}
				if len(changed) == 0 {
					g.Log().Infof(ctx, "no changed products")
					return
				}

				in := &model.NotifyInput{}
				for i, p := range changed {
					in.Text += fmt.Sprintf("name: %s, has: %t", p.Name, !p.SoldOut)
					if i != len(changed)-1 {
						in.Text += "\n"
					}
				}

				err = service.Content().Notify(ctx, in)
				if err != nil {
					g.Log().Errorf(ctx, "call notify err: %s", err)
					return
				}
			})
			if err != nil {
				return err
			}

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Info(ctx, `cron job shutdown`)
			})
			g.Listen()
			return
		},
	}
)
