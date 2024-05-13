package unix

import (
	"time"

	"github.com/kitanoyoru/kitad/pkg/types"
)

type Option func(*Options)

type Options struct {
	DialTimeout         *time.Duration
	KeepAlive           *time.Duration
	TLSHandshakeTimeout *time.Duration
}

func WithDialTimeout(d time.Duration) Option {
	return func(opts *Options) {
		opts.DialTimeout = types.Duration(d)
	}
}

func WithKeepAlive(d time.Duration) Option {
	return func(opts *Options) {
		opts.KeepAlive = types.Duration(d)
	}
}

func WithTLSHandshakeTimeout(d time.Duration) Option {
	return func(opts *Options) {
		opts.TLSHandshakeTimeout = types.Duration(d)
	}
}
