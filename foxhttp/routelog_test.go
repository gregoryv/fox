package foxhttp

import (
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
)

func TestRouteLog_middleware(t *testing.T) {
	log := NewRouteLog(t)
	var called bool
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusBadRequest)
	})
	handler := log.Middleware(next)
	assert := asserter.New(t)
	exp := assert().ResponseFrom(handler)

	exp.StatusCode(400, "GET", "/")
	assert(called).Error("next not called")
}
