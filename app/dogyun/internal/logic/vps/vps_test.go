package vps

import (
	"context"
	"fmt"
	"testing"

	"my-crontab/app/dogyun/internal/model"
)

func TestSContent_HasVPS(t *testing.T) {
	s := &sContent{}
	res, err := s.HasVPS(context.Background(), &model.HasVPSInput{
		ProductGroup: 10,
		Name:         "jp.iij.amd.s",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("res: %+v\n", res)
}

func TestSContent_Notify(t *testing.T) {
	s := &sContent{}
	err := s.Notify(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
