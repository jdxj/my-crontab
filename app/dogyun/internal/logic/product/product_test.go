package product

import (
	"context"
	"testing"
)

func TestSContent_Notify(t *testing.T) {
	s := &sContent{}
	err := s.Notify(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
