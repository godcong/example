package main

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	logger "github.com/origadmin/slog-kratos"
	"github.com/origadmin/toolkits/codec"
	ccfg "github.com/origadmin/toolkits/contrib/config"

	"example/helloworld/helloworld"
)

type Flags struct {
	ID         string
	Name       string
	Version    string
	EnvPath    string
	ConfigPath string
	MetaData   map[string]string
}

var flags Flags

func init() {
	flags = Flags{
		ID:         "helloworld",
		Name:       "helloworld",
		Version:    "latest",
		EnvPath:    ".env",
		ConfigPath: "config",
		MetaData:   map[string]string{},
	}
}

func initLogger() log.Logger {
	l := log.With(logger.NewLogger(),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", flags.ID,
		"service.name", flags.Name,
		"service.version", flags.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(l)
	return l
}

func main() {
	initLogger()

	s := file.NewSource("configs")

	cfg := config.New(config.WithSource(s), config.WithDecoder(ccfg.SourceDecoder))
	if err := cfg.Load(); err != nil {
		panic(err)
	}
	var bs helloworld.Bootstrap
	if err := cfg.Scan(&bs); err != nil {
		panic(err)
	}
	log.Infof("service default config: %v", bs.ServiceName)

	if err := codec.DecodeFromFile("configs/example.toml", &bs); err != nil {
		panic(err)
	}
	log.Infof("service toml config: %v", bs.ServiceName)

	if err := codec.DecodeFromFile("configs/example.yml", &bs); err != nil {
		panic(err)
	}
	log.Infof("service yaml config: %v", bs.ServiceName)
}
