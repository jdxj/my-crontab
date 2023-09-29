package service

import (
	"context"

	"my-crontab/app/dogyun/internal/model"
)

type (
	IContent interface {
		HasVPS(ctx context.Context, in *model.HasVPSInput) (*model.HasVPSOutput, error)
		Notify(ctx context.Context, in *model.NotifyInput) error
	}
)

var (
	localContent IContent
)

func Content() IContent {
	if localContent == nil {
		panic("implement not found for interface IContent, forgot register?")
	}
	return localContent
}

func RegisterContent(i IContent) {
	localContent = i
}
