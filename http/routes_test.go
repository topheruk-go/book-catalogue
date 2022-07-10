package http

import (
	"net/http/httptest"
	"testing"

	test "github.com/topheruk-go/util/testing/http"
)

func TestRoutes(t *testing.T) {
	var ser Service //TODO
	test.HttpService(t, httptest.NewServer(ser), []test.HttpCase{})
}
