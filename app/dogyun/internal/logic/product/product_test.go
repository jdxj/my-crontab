package product

import (
	"context"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

func TestSContent_Notify(t *testing.T) {
	s := &sContent{}
	err := s.Notify(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestSContent_GetProducts(t *testing.T) {
	g.Log().SetLevel(glog.LEVEL_DEBU)
	res, err := getProducts(context.Background(), g.Client(), 13)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Println(res)
}
