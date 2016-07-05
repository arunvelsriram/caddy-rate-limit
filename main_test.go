package ratelimit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	server *httptest.Server
	cl     *CustomLimiter
)

func TestMain(m *testing.M) {

	fakeHandler := func(w http.ResponseWriter, r *http.Request) {
		ip := GetRemoteIP(r)
		fmt.Fprint(w, ip)
	}
	server = httptest.NewServer(http.HandlerFunc(fakeHandler))
	defer server.Close()

	cl = NewCustomLimiter()

	os.Exit(m.Run())
}
