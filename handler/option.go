package handler

import (
	"crypto/tls"
	"net/url"

	"github.com/liukeqqs/core/auth"
	"github.com/liukeqqs/core/bypass"
	"github.com/liukeqqs/core/chain"
	"github.com/liukeqqs/core/limiter/rate"
	"github.com/liukeqqs/core/limiter/traffic"
	"github.com/liukeqqs/core/logger"
	"github.com/liukeqqs/core/metadata"
	"github.com/liukeqqs/core/observer"
)

type Options struct {
	Bypass      bypass.Bypass
	Router      *chain.Router
	Auth        *url.Userinfo
	Auther      auth.Authenticator
	RateLimiter rate.RateLimiter
	Limiter     traffic.TrafficLimiter
	TLSConfig   *tls.Config
	Logger      logger.Logger
	Observer    observer.Observer
	Service     string
}

type Option func(opts *Options)

func BypassOption(bypass bypass.Bypass) Option {
	return func(opts *Options) {
		opts.Bypass = bypass
	}
}

func RouterOption(router *chain.Router) Option {
	return func(opts *Options) {
		opts.Router = router
	}
}

func AuthOption(auth *url.Userinfo) Option {
	return func(opts *Options) {
		opts.Auth = auth
	}
}

func AutherOption(auther auth.Authenticator) Option {
	return func(opts *Options) {
		opts.Auther = auther
	}
}

func RateLimiterOption(limiter rate.RateLimiter) Option {
	return func(opts *Options) {
		opts.RateLimiter = limiter
	}
}

func TrafficLimiterOption(limiter traffic.TrafficLimiter) Option {
	return func(opts *Options) {
		opts.Limiter = limiter
	}
}

func TLSConfigOption(tlsConfig *tls.Config) Option {
	return func(opts *Options) {
		opts.TLSConfig = tlsConfig
	}
}

func LoggerOption(logger logger.Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}

func ObserverOption(observer observer.Observer) Option {
	return func(opts *Options) {
		opts.Observer = observer
	}
}

func ServiceOption(service string) Option {
	return func(opts *Options) {
		opts.Service = service
	}
}

type HandleOptions struct {
	Metadata metadata.Metadata
}

type HandleOption func(opts *HandleOptions)

func MetadataHandleOption(md metadata.Metadata) HandleOption {
	return func(opts *HandleOptions) {
		opts.Metadata = md
	}
}
