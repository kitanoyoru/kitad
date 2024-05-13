package unix

import (
	"time"
)

type Config struct {
	DialTimeout         time.Duration
	KeepAlive           time.Duration
	TLSHandshakeTimeout time.Duration
}

func NewConfigFromOpts(opts *Options) *Config {
	c := Config{
		DialTimeout:         30 * time.Second,
		KeepAlive:           30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	if opts.DialTimeout != nil {
		c.DialTimeout = *opts.DialTimeout
	}

	if opts.KeepAlive != nil {
		c.KeepAlive = *opts.KeepAlive
	}

	if opts.TLSHandshakeTimeout != nil {
		c.TLSHandshakeTimeout = *opts.TLSHandshakeTimeout
	}

	return &c
}
