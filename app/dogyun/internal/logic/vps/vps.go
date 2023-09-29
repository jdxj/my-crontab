package vps

import (
	"context"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"my-crontab/app/dogyun/internal/consts"
	"my-crontab/app/dogyun/internal/model"
	"my-crontab/app/dogyun/internal/service"
)

func init() {
	service.RegisterContent(&sContent{})
}

type sContent struct{}

func (s *sContent) HasVPS(ctx context.Context, in *model.HasVPSInput) (*model.HasVPSOutput, error) {
	cookie, err := g.Cfg().Get(ctx, "dogyun.cookie")
	if err != nil {
		return nil, err
	}
	xct, err := g.Cfg().Get(ctx, "dogyun.x-csrf-token")
	if err != nil {
		return nil, err
	}

	httpClient := g.Client().
		SetHeaderMap(consts.Header).
		SetHeader("cookie", cookie.String()).
		SetHeader("x-csrf-token", xct.String())

	res, err := httpClient.PostForm(ctx, consts.Target, map[string]string{
		"productGroup": gconv.String(in.ProductGroup),
	})
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var m []*model.VPS
	err = json.Unmarshal(res.ReadAll(), &m)
	if err != nil {
		return nil, err
	}

	out := &model.HasVPSOutput{}
	for _, v := range m {
		if v.Name == in.Name && !v.SoldOut {
			out.Has = true
		}
	}
	return out, nil
}

func (s *sContent) Notify(ctx context.Context, in *model.NotifyInput) error {
	chatId, err := g.Cfg().Get(ctx, "tg.chatId")
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatId.Int64(), in.Text)

	token, err := g.Cfg().Get(ctx, "tg.token")
	if err != nil {
		return err
	}
	bot, err := tgbotapi.NewBotAPI(token.String())
	if err != nil {
		return err
	}
	_, err = bot.Send(msg)
	return err
}
