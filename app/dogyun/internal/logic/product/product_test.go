package product

import (
	"context"
	"fmt"
	"testing"
)

func TestSContent_Notify(t *testing.T) {
	s := &sContent{}
	err := s.Notify(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestSContent_GetProducts(t *testing.T) {
	s := newSContent()
	res, err := s.GetProducts(context.Background())
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}
