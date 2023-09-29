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
		Arguments: []gcmd.Argument{
			{
				Name:  "product-group",
				Short: "p",
			},
			{
				Name:  "name",
				Short: "n",
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, `cron job start`)

			every, err := g.Cfg().Get(ctx, "dogyun.every")
			if err != nil {
				return err
			}

			in := &model.HasVPSInput{
				ProductGroup: parser.GetOpt("p").Int(),
				Name:         parser.GetOpt("n").String(),
			}

			pattern := fmt.Sprintf("@every %s", every.String())
			_, err = gcron.Add(ctx, pattern, func(ctx context.Context) {
				res, err := service.Content().HasVPS(ctx, in)
				if err != nil {
					g.Log().Errorf(ctx, "call hasVPS err: %s", err)
					return
				}

				if !res.Has {
					return
				}

				err = service.Content().Notify(ctx, &model.NotifyInput{
					Text: fmt.Sprintf("%s中了", in.Name),
				})
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
