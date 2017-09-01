package router

import (
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()

	if router.Get("Free rooms") == nil {
		t.Error("Something wrong with routers")
	}

	if router.Get("Free rooms1") != nil {
		t.Error("Something wrong with routers")
	}
}
