package foxhttp

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/fox"
)

func TestRouteLog_middleware(t *testing.T) {
	var (
		buf  bytes.Buffer
		out  = fox.NewSyncLog(&buf)
		log  = NewRouteLog(out)
		next = func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}
		handler = log.MiddlewareFunc(next)
	)

	assert := asserter.New(t)
	exp := assert().ResponseFrom(handler)

	exp.StatusCode(400, "GET", "/")
	assert(buf.String()[:5] == "GET /").Error(buf.String())
}
