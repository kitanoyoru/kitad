package unix

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/kitanoyoru/kitad/pkg/transport"
)

func NewTransport(req ConnectionRequest, opts ...Option) (transport.Transport, error) {
	transportOpts := Options{}

	for _, o := range opts {
		o(&transportOpts)
	}

	config := NewConfigFromOpts(&transportOpts)

	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   config.DialTimeout,
			KeepAlive: config.KeepAlive,
		}).DialContext,
		TLSHandshakeTimeout: config.TLSHandshakeTimeout,
		// TODO: add tls client config
	}

	dialer := &net.Dialer{
		Timeout:   config.DialTimeout,
		KeepAlive: config.KeepAlive,
	}

	dialContext := func(ctx context.Context, net, addr string) (net.Conn, error) {
		return dialer.DialContext(ctx, "unix", addr)
	}

	ut := &unixTransport{
		&http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			DialContext:         dialContext,
			TLSHandshakeTimeout: config.TLSHandshakeTimeout,
			// TODO: add tls client config here
			IdleConnTimeout: time.Microsecond,
		},
	}

	t.RegisterProtocol("unix", ut)
	t.RegisterProtocol("unixs", ut)

	return t, nil
}

type unixTransport struct {
	*http.Transport
}

func (ut *unixTransport) RoundTripper(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = strings.Replace(req.URL.Scheme, "unix", "http", 1)
	return ut.Transport.RoundTrip(req)
}
