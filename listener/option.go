package listener

import (
	"crypto/tls"
	"net/url"

	"github.com/liukeqqs/core/admission"
	"github.com/liukeqqs/core/auth"
	"github.com/liukeqqs/core/chain"
	"github.com/liukeqqs/core/limiter/conn"
	"github.com/liukeqqs/core/limiter/traffic"
	"github.com/liukeqqs/core/logger"
	"github.com/liukeqqs/x/stats"
)

type Options struct {
	Addr           string
	Auther         auth.Authenticator
	Auth           *url.Userinfo
	TLSConfig      *tls.Config
	Admission      admission.Admission
	TrafficLimiter traffic.TrafficLimiter
	ConnLimiter    conn.ConnLimiter
	Chain          chain.Chainer
	Stats          *stats.Stats
	Logger         logger.Logger
	Service        string
	ProxyProtocol  int
}

type Option func(opts *Options)

func AddrOption(addr string) Option {
	return func(opts *Options) {
		opts.Addr = addr
	}
}

func AutherOption(auther auth.Authenticator) Option {
	return func(opts *Options) {
		opts.Auther = auther
	}
}

func AuthOption(auth *url.Userinfo) Option {
	return func(opts *Options) {
		opts.Auth = auth
	}
}

func TLSConfigOption(tlsConfig *tls.Config) Option {
	return func(opts *Options) {
		opts.TLSConfig = tlsConfig
	}
}

func AdmissionOption(admission admission.Admission) Option {
	return func(opts *Options) {
		opts.Admission = admission
	}
}

func TrafficLimiterOption(limiter traffic.TrafficLimiter) Option {
	return func(opts *Options) {
		opts.TrafficLimiter = limiter
	}
}

func ConnLimiterOption(limiter conn.ConnLimiter) Option {
	return func(opts *Options) {
		opts.ConnLimiter = limiter
	}
}

func ChainOption(chain chain.Chainer) Option {
	return func(opts *Options) {
		opts.Chain = chain
	}
}

func StatsOption(stats *stats.Stats) Option {
	return func(opts *Options) {
		opts.Stats = stats
	}
}

func LoggerOption(logger logger.Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}

func ServiceOption(service string) Option {
	return func(opts *Options) {
		opts.Service = service
	}
}

func ProxyProtocolOption(ppv int) Option {
	return func(opts *Options) {
		opts.ProxyProtocol = ppv
	}
}
