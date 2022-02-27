package elastic_encoder

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/logging"
	"go.uber.org/zap/zapcore"
)

func init() {
	caddy.RegisterModule(ElasticEncoder{})
}

type ElasticEncoder struct {
	logging.LogEncoderConfig
	zapcore.Encoder `json:"-"`
}

func (ElasticEncoder) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "caddy.logging.encoders.elastic",
		New: func() caddy.Module {
			return new(ElasticEncoder)
		},
	}
}

func (e *ElasticEncoder) Provision(_ caddy.Context) error {
	cfg := e.ZapcoreEncoderConfig()
	cfg.LevelKey = "log.level"
	cfg.TimeKey = "@timestamp"
	cfg.NameKey = "log.logger"
	cfg.MessageKey = "message"
	cfg.CallerKey = "log.caller"
	cfg.FunctionKey = "log.origin.function"
	cfg.StacktraceKey = "error.stack_trace"
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	e.Encoder = zapcore.NewJSONEncoder(cfg)

	return nil
}
