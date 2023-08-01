package gameofthrones

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	log := fmt.Sprintf("[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	fmt.Fprint(l.logger, log)
	return l.next.RoundTrip(r)
}
