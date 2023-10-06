package product

import (
	"context"
	"encoding/json"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"

	"my-crontab/app/dogyun/internal/consts"
	"my-crontab/app/dogyun/internal/model"
	"my-crontab/app/dogyun/internal/service"
)

func init() {
	service.RegisterContent(newSContent())
}

func newSContent() *sContent {
	return &sContent{
		stat: make(map[string]bool),
	}
}

type sContent struct {
	// name: soldOut
	stat map[string]bool
}

func (s *sContent) GetProducts(ctx context.Context) ([]*model.Product, error) {
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
		SetHeader("x-csrf-token", xct.String()).
		Timeout(time.Second * 30)

	var res []*model.Product
	for _, pg := range consts.ProductGroups {
		p, err := getProducts(ctx, httpClient, pg)
		if err != nil {
			return nil, err
		}

		res = append(res, p...)
	}
	return res, nil
}

func getProducts(ctx context.Context, c *gclient.Client, productGroup int) ([]*model.Product, error) {
	if productGroup <= 0 {
		g.Log().Warningf(ctx, "invalid productGroup: %d", productGroup)
		return nil, nil
	}

	res, err := c.Post(ctx, consts.Target, g.Map{
		"productGroup": productGroup,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		g.Log().Debugf(ctx, "req/res info: %s", res.Raw())

		err := res.Close()
		if err != nil {
			g.Log().Warningf(ctx, "close http response err: %s", err)
		}
	}()

	var m []*model.Product
	return m, json.Unmarshal(res.ReadAll(), &m)
}

func (s *sContent) GetChangedProducts(ctx context.Context) ([]*model.Product, error) {
	ps, err := s.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	var changedProducts []*model.Product
	for _, p := range ps {
		if s.stat[p.Name] == p.SoldOut {
			continue
		}

		s.stat[p.Name] = p.SoldOut
		changedProducts = append(changedProducts, p)
	}
	return changedProducts, nil
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
