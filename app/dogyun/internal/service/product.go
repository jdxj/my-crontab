// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"my-crontab/app/dogyun/internal/model"
)

type (
	IContent interface {
		GetProducts(ctx context.Context) ([]*model.Product, error)
		GetChangedProducts(ctx context.Context) ([]*model.Product, error)
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
